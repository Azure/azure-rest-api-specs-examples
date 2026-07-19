
/**
 * Samples for DataWarehouseUserActivitiesOperation ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListDataWarehouseUserActivities.json
     */
    /**
     * Sample code: List of the user activities of a data warehouse.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listOfTheUserActivitiesOfADataWarehouse(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataWarehouseUserActivitiesOperations().listByDatabase("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", com.azure.core.util.Context.NONE);
    }
}
