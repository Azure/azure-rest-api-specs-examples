
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceLongTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.BackupStorageAccessTier;
import com.azure.resourcemanager.sql.models.ManagedInstanceLongTermRetentionPolicyName;

/**
 * Samples for ManagedInstanceLongTermRetentionPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceLongTermRetentionPolicyCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update the LTR policy for the managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createOrUpdateTheLTRPolicyForTheManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceLongTermRetentionPolicies().createOrUpdate("testResourceGroup",
            "testInstance", "testDatabase", ManagedInstanceLongTermRetentionPolicyName.DEFAULT,
            new ManagedInstanceLongTermRetentionPolicyInner().withBackupStorageAccessTier(BackupStorageAccessTier.HOT)
                .withWeeklyRetention("P1M").withMonthlyRetention("P1Y").withYearlyRetention("P5Y").withWeekOfYear(5),
            com.azure.core.util.Context.NONE);
    }
}
