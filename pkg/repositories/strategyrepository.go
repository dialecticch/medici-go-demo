package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/dialecticch/medici-go/pkg/types"
)

type onUpdateFunc func(s []*types.StrategyConfig)

type ActiveFlag int

const (
	Inactive ActiveFlag = iota
	Active
	Any
)

type StrategyRepository struct {
	mux sync.RWMutex

	db         *sql.DB
	strategies []*types.StrategyConfig

	onUpdateFunc onUpdateFunc

	chainID uint64
}

func NewStrategyRepository(db *sql.DB, chainID uint64) *StrategyRepository {
	return &StrategyRepository{
		mux:     sync.RWMutex{},
		db:      db,
		chainID: chainID,
	}
}

func (s *StrategyRepository) Run(flag ActiveFlag) error {
	strategies, err := s.Query(flag)
	if err != nil {
		return err
	}

	s.strategies = strategies

	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			strategies, err := s.Query(flag)
			if err != nil {
				log.Printf("failed to query strategies err: %s", err)
				continue
			}

			s.mux.Lock()
			s.strategies = strategies
			s.mux.Unlock()

			s.mux.RLock()
			f := s.onUpdateFunc
			s.mux.RUnlock()

			if f != nil {
				s.onUpdateFunc(strategies)
			}
		}
	}()

	return nil
}

func (s *StrategyRepository) Map(f func(s *types.StrategyConfig)) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	for _, s := range s.strategies {
		f(s)
	}
}

func (s *StrategyRepository) OnUpdate(f onUpdateFunc) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.onUpdateFunc = f
}

func (s *StrategyRepository) Query(flag ActiveFlag) ([]*types.StrategyConfig, error) {

	args := []interface{}{s.chainID}

	query := `
		SELECT
	   		safes.address,
			latest_versions.address,
			strategies.name,
		   	safe_pools.threshold,
		   	safe_pools.action,
		   	safe_pools.threshold_type,
		   	safe_pools.last_harvested,
		   	pools.pool,
		   	tokens.address
		FROM safe_pools
			INNER JOIN safes
				ON safe_pools.safe_id = safes.id
			INNER JOIN pools
				ON safe_pools.pool_id = pools.id
			INNER JOIN latest_versions 
				ON pools.strategy_id = latest_versions.strategy_id
			INNER JOIN strategies 
				ON pools.strategy_id = strategies.id
			INNER JOIN tokens 
				ON pools.token_id = tokens.id
			WHERE safes.chain_id = $1`

	if flag != Any {
		query += " AND safe_pools.active = $2"
		args = append(args, flag == Active)
	}

	rows, err := s.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	data := make([]*types.StrategyConfig, 0)

	for rows.Next() {
		var safe string
		var strategy string
		var name string
		var threshold float64
		var action string
		var thresholdType string
		var lastHarvested string
		var rawpool string
		var depositToken string

		err := rows.Scan(&safe, &strategy, &name, &threshold, &action, &thresholdType, &lastHarvested, &rawpool, &depositToken)
		if err != nil {
			return nil, err
		}

		block, ok := new(big.Int).SetString(lastHarvested, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse block %s", lastHarvested)
		}

		pool, ok := new(big.Int).SetString(rawpool, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse pool %s", rawpool)
		}

		data = append(data, &types.StrategyConfig{
			Safe:         common.HexToAddress(safe),
			DepositToken: common.HexToAddress(depositToken),
			Strategy: types.StrategyMetadata{
				Name:    name,
				Address: common.HexToAddress(strategy),
			},
			Threshold:     threshold,
			Action:        action,
			ThresholdType: thresholdType,
			LastHarvested: *block,
			Pool:          pool,
		})
	}

	return data, nil
}
