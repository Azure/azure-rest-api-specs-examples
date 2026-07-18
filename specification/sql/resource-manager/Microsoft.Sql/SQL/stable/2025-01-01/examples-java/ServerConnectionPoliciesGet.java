
import com.azure.resourcemanager.sql.models.ConnectionPolicyName;

/**
 * Samples for ServerConnectionPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerConnectionPoliciesGet.json
     */
    /**
     * Sample code: Gets a server connection policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAServerConnectionPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerConnectionPolicies().getWithResponse("rgtest-12", "servertest-6285",
            ConnectionPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
