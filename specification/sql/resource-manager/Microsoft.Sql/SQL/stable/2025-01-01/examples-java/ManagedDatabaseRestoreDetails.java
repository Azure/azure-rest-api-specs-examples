
import com.azure.resourcemanager.sql.models.RestoreDetailsName;

/**
 * Samples for ManagedDatabaseRestoreDetails Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseRestoreDetails.json
     */
    /**
     * Sample code: Managed database restore details.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void managedDatabaseRestoreDetails(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseRestoreDetails().getWithResponse("Default-SQL-SouthEastAsia",
            "managedInstance", "testdb", RestoreDetailsName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
