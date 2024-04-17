
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/** Samples for ManagedBackupShortTermRetentionPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/UpdateManagedShortTermRetentionPolicy
     * .json
     */
    /**
     * Sample code: Update the short term retention policy for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedBackupShortTermRetentionPolicies().createOrUpdate(
            "resourceGroup", "testsvr", "testdb", ManagedShortTermRetentionPolicyName.DEFAULT,
            new ManagedBackupShortTermRetentionPolicyInner().withRetentionDays(14), Context.NONE);
    }
}
