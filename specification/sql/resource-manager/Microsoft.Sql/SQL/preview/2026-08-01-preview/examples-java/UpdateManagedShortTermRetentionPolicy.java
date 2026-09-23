
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/**
 * Samples for ManagedBackupShortTermRetentionPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/UpdateManagedShortTermRetentionPolicy.json
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
            new ManagedBackupShortTermRetentionPolicyInner().withRetentionDays(14).withLockImmutability(false),
            com.azure.core.util.Context.NONE);
    }
}
