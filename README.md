# medici-go

This repository contains the execution for the Medici on-chain system. The following tools are included:
 - **Control:** These are all tools to manage automation and reporting.
   - **Automator:** The automator consistently checks all strategies and harvests when a certain criteria was reached as defined in the database.
   - **Balances:** Monitors and logs any balance updates to the strategies.
   - **Harvests:** Monitors and logs any harvests.
   - **Safe:** Monitors and logs any safe changes.
 - **Admin:** These are all the admin tools required:
   - **Backfill:** In case we missed any harvests or balance updates we can manually import them to the database.
   - **Tokens:** Adds a tokens information to the database based on its address.
   - **Strategies:**
     - **List:** Lists the strategies in the database
     - **Add:** Adds a new strategy to the database.
   - **Pools:**
      - **Add:** Imports a strategies specific pool info.
      - **Connect:** Connects a pool to a safe
