
/**
 * Samples for Databases UpgradeDataWarehouse.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/UpgradeDataWarehouse.json
     */
    /**
     * Sample code: Upgrades a data warehouse.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void upgradesADataWarehouse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().upgradeDataWarehouse("Default-SQL-SouthEastAsia",
            "testsvr", "testdwdb", com.azure.core.util.Context.NONE);
    }
}
