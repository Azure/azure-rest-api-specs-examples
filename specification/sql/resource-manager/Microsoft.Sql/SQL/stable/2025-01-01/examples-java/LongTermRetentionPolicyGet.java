
import com.azure.resourcemanager.sql.models.LongTermRetentionPolicyName;

/**
 * Samples for LongTermRetentionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LongTermRetentionPolicyGet.json
     */
    /**
     * Sample code: Get the long term retention policy for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheLongTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionPolicies().getWithResponse("resourceGroup", "testserver",
            "testDatabase", LongTermRetentionPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
