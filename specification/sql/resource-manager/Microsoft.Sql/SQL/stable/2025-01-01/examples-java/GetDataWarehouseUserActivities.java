
import com.azure.resourcemanager.sql.models.DataWarehouseUserActivityName;

/**
 * Samples for DataWarehouseUserActivitiesOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetDataWarehouseUserActivities.json
     */
    /**
     * Sample code: Get the list of the user activities of a data warehouse.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheListOfTheUserActivitiesOfADataWarehouse(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataWarehouseUserActivitiesOperations().getWithResponse("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", DataWarehouseUserActivityName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
