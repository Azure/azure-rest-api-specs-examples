import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ManagedBackupShortTermRetentionPolicyInner;
import com.azure.resourcemanager.sql.models.ManagedShortTermRetentionPolicyName;

/** Samples for ManagedBackupShortTermRetentionPolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/UpdateManagedShortTermRetentionPolicy.json
     */
    /**
     * Sample code: Update the short term retention policy for the database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTheShortTermRetentionPolicyForTheDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedBackupShortTermRetentionPolicies()
            .update(
                "resourceGroup",
                "testsvr",
                "testdb",
                ManagedShortTermRetentionPolicyName.DEFAULT,
                new ManagedBackupShortTermRetentionPolicyInner(),
                Context.NONE);
    }
}
