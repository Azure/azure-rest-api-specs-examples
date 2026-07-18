
import com.azure.resourcemanager.sql.models.ManagedInstanceLongTermRetentionPolicyName;

/**
 * Samples for ManagedInstanceLongTermRetentionPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceLongTermRetentionPolicyDelete.json
     */
    /**
     * Sample code: Deletes the LTR policy for the managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesTheLTRPolicyForTheManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceLongTermRetentionPolicies().delete("testResourceGroup",
            "testInstance", "testDatabase", ManagedInstanceLongTermRetentionPolicyName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
