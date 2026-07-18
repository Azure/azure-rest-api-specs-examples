
import com.azure.resourcemanager.sql.models.ShortTermRetentionPolicyName;

/**
 * Samples for BackupShortTermRetentionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetShortTermRetentionPolicy.json
     */
    /**
     * Sample code: Get the short term retention policy for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getBackupShortTermRetentionPolicies().getWithResponse("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", ShortTermRetentionPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
