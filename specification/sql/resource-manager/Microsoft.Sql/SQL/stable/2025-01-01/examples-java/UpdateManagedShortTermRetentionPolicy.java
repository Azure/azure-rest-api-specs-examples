
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/**
 * Samples for ManagedBackupShortTermRetentionPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/UpdateManagedShortTermRetentionPolicy.json
     */
    /**
     * Sample code: Update the short term retention policy for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedBackupShortTermRetentionPolicies().update("resourceGroup", "testsvr",
            "testdb", ManagedShortTermRetentionPolicyName.DEFAULT,
            new ManagedBackupShortTermRetentionPolicyInner().withRetentionDays(14), com.azure.core.util.Context.NONE);
    }
}
