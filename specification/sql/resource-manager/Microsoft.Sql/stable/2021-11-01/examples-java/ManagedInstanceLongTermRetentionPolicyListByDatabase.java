
import com.azure.core.util.Context;

/** Samples for ManagedInstanceLongTermRetentionPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceLongTermRetentionPolicyListByDatabase.json
     */
    /**
     * Sample code: Get the long term retention policies for the managed database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheLongTermRetentionPoliciesForTheManagedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceLongTermRetentionPolicies()
            .listByDatabase("testResourceGroup", "testInstance", "testDatabase", Context.NONE);
    }
}
