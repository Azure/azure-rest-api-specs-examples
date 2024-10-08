
import com.azure.resourcemanager.sql.models.DataWarehouseUserActivityName;

/**
 * Samples for DataWarehouseUserActivitiesOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetDataWarehouseUserActivities.json
     */
    /**
     * Sample code: Get the list of the user activities of a data warehouse.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheListOfTheUserActivitiesOfADataWarehouse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDataWarehouseUserActivitiesOperations().getWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb", DataWarehouseUserActivityName.CURRENT,
            com.azure.core.util.Context.NONE);
    }
}
