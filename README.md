# medici-go

The Medici is infrastructure used by Dialectic to automate its yield farming operation. This repository contains the backend to provide functionality
required by the Dialectic team.

For further details check out the following articles:
 - [Medici: Dialectic’s Yield Farmer
](https://dialectic.ch/editorial/medici-dialectics-yield-farmer)
 - [Chronograph: Institutional Grade Yield Farming
](https://dialectic.ch/editorial/chronograph-overview)

Contracts can be found [here](https://github.com/dialecticch/medici-contracts-demo).

## Tools

Medici is made up of several tools.

### Control `ctrl`

These are the tools used to manage automation and reporting.

- **Automator `automate`:** The automator consistently checks all strategies and harvests when a certain criteria was reached as defined in the database.
- **Balances `balances`:** Monitors and logs any balance updates to the strategies.
- **Harvests `harvests`:** Monitors and logs any harvests.
- **Safe `safe`:** Monitors and logs any safe changes.

### Admin `admin`

These are the admin tools mainly used for database management

- **Admin:** These are all the admin tools required:
  - **Backfill `backfill`:** In case we missed any harvests or balance updates we can manually import them to the database.
  - **Cleaner `cleaner`:** Checks for dust balances in strategies and allows claiming them.
  - **Tokens `tokens`**
    - **Add `add`:** Adds a tokens information to the database based on its address.
  - **Strategies `strategies`:**
    - **List `list`:** Lists the strategies in the database
    - **Add `add`:** Adds a new strategy to the database.
  - **Pools `pools`:**
     - **Add `add`:** Imports a strategies specific pool info.
     - **Connect `connect`:** Connects a pool to a safe
