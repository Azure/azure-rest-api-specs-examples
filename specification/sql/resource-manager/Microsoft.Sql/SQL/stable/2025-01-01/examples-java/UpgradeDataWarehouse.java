
/**
 * Samples for Databases UpgradeDataWarehouse.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/UpgradeDataWarehouse.json
     */
    /**
     * Sample code: Upgrades a data warehouse.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void upgradesADataWarehouse(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().upgradeDataWarehouse("Default-SQL-SouthEastAsia", "testsvr", "testdwdb",
            com.azure.core.util.Context.NONE);
    }
}
