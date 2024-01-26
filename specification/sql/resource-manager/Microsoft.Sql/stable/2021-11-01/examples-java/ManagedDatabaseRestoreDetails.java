
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.RestoreDetailsName;

/** Samples for ManagedDatabaseRestoreDetails Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseRestoreDetails.json
     */
    /**
     * Sample code: Managed database restore details.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void managedDatabaseRestoreDetails(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseRestoreDetails().getWithResponse(
            "Default-SQL-SouthEastAsia", "managedInstance", "testdb", RestoreDetailsName.DEFAULT, Context.NONE);
    }
}
