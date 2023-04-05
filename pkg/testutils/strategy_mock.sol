// SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.2;

contract MockStrategy {
    struct Harvest {
        address token;
        uint256 amount;
    }

    event Harvested(
        uint256 pool,
        address indexed safe,
        address indexed token,
        uint256 amount
    );
    event Deposited(uint256 pool, address indexed safe, uint256 amount);
    event Withdrew(uint256 pool, address indexed safe, uint256 amount);

    function logHarvested(
        uint256 pool,
        address safe,
        address token,
        uint256 amount
    ) public {
        emit Harvested(pool, safe, token, amount);
    }

    function logDeposited(
        uint256 pool,
        address safe,
        uint256 amount
    ) public {
        emit Deposited(pool, safe, amount);
    }

    function logWithdrew(
        uint256 pool,
        address safe,
        uint256 amount
    ) public {
        emit Withdrew(pool, safe, amount);
    }

    function harvest(
        uint256 pool,
        address _safe,
        bytes calldata data
    ) public {
        1 + 1;
    }

    function compound(
        uint256 pool,
        address _safe,
        bytes calldata data
    ) public {
        1 + 1;
    }

    function NAME() public returns (string memory) {
        return "Mock Strategy";
    }

    function depositedAmount(uint256 pool, address safe)
        external
        view
        returns (uint256)
    {
        return 10;
    }

    function depositToken() external view returns (address) {
        return address(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48);
    }

    function simulateClaim(
        uint256 pool,
        address safe,
        bytes calldata data
    ) public returns (Harvest[] memory) {
        Harvest[] memory harvests = new Harvest[](2);
        harvests[0] = Harvest(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48, 10);
        harvests[1] = Harvest(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2, 10);
        return harvests;
    }
}
