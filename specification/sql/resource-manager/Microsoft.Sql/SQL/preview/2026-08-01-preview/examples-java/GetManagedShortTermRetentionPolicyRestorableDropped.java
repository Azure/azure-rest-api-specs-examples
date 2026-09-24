
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/**
 * Samples for ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetManagedShortTermRetentionPolicyRestorableDropped.json
     */
    /**
     * Sample code: Get the short term retention policy for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies().getWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000",
            ManagedShortTermRetentionPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
