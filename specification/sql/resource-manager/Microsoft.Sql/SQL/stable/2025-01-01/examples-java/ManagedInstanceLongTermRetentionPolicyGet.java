
import com.azure.resourcemanager.sql.models.ManagedInstanceLongTermRetentionPolicyName;

/**
 * Samples for ManagedInstanceLongTermRetentionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceLongTermRetentionPolicyGet.json
     */
    /**
     * Sample code: Get the long term retention policy for the managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheLongTermRetentionPolicyForTheManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceLongTermRetentionPolicies().getWithResponse("testResourceGroup",
            "testInstance", "testDatabase", ManagedInstanceLongTermRetentionPolicyName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
