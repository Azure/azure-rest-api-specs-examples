
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/**
 * Samples for ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/UpdateManagedShortTermRetentionPolicyRestorableDropped.json
     */
    /**
     * Sample code: Update the short term retention policy for the restorable dropped database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateTheShortTermRetentionPolicyForTheRestorableDroppedDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies().update(
            "resourceGroup", "testsvr", "testdb,131403269876900000", ManagedShortTermRetentionPolicyName.DEFAULT,
            new ManagedBackupShortTermRetentionPolicyInner().withRetentionDays(14).withLockImmutability(false),
            com.azure.core.util.Context.NONE);
    }
}
