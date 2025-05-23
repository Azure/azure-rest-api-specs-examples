
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/**
 * Samples for ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * UpdateManagedShortTermRetentionPolicyRestorableDropped.json
     */
    /**
     * Sample code: Update the short term retention policy for the restorable dropped database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheShortTermRetentionPolicyForTheRestorableDroppedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient()
            .getManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies().update("resourceGroup", "testsvr",
                "testdb,131403269876900000", ManagedShortTermRetentionPolicyName.DEFAULT,
                new ManagedBackupShortTermRetentionPolicyInner().withRetentionDays(14),
                com.azure.core.util.Context.NONE);
    }
}
