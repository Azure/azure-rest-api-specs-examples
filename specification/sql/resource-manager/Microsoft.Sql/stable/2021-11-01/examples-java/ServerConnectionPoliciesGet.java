
import com.azure.resourcemanager.sql.models.ConnectionPolicyName;

/**
 * Samples for ServerConnectionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerConnectionPoliciesGet.json
     */
    /**
     * Sample code: Gets a server connection policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAServerConnectionPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerConnectionPolicies().getWithResponse("rgtest-12",
            "servertest-6285", ConnectionPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
