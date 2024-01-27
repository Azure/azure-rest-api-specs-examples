
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.ManagedInstanceLongTermRetentionPolicyName;

/** Samples for ManagedInstanceLongTermRetentionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceLongTermRetentionPolicyGet.json
     */
    /**
     * Sample code: Get the long term retention policy for the managed database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheLongTermRetentionPolicyForTheManagedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceLongTermRetentionPolicies().getWithResponse(
            "testResourceGroup", "testInstance", "testDatabase", ManagedInstanceLongTermRetentionPolicyName.DEFAULT,
            Context.NONE);
    }
}
