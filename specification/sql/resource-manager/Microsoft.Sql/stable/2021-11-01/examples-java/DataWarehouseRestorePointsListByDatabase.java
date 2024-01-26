
import com.azure.core.util.Context;

/** Samples for RestorePoints ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * DataWarehouseRestorePointsListByDatabase.json
     */
    /**
     * Sample code: List datawarehouse database restore points.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatawarehouseDatabaseRestorePoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRestorePoints().listByDatabase("Default-SQL-SouthEastAsia",
            "testserver", "testDatabase", Context.NONE);
    }
}
