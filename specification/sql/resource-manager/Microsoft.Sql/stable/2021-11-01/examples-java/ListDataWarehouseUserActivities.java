
/**
 * Samples for DataWarehouseUserActivitiesOperation ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListDataWarehouseUserActivities.json
     */
    /**
     * Sample code: List of the user activities of a data warehouse.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listOfTheUserActivitiesOfADataWarehouse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDataWarehouseUserActivitiesOperations()
            .listByDatabase("Default-SQL-SouthEastAsia", "testsvr", "testdb", com.azure.core.util.Context.NONE);
    }
}
