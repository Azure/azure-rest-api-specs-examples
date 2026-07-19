
import com.azure.resourcemanager.sql.fluent.models.ServerConnectionPolicyInner;
import com.azure.resourcemanager.sql.models.ConnectionPolicyName;
import com.azure.resourcemanager.sql.models.ServerConnectionType;

/**
 * Samples for ServerConnectionPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerConnectionPoliciesUpdate.json
     */
    /**
     * Sample code: Updates a server connection policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updatesAServerConnectionPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerConnectionPolicies().createOrUpdate("testrg", "testserver",
            ConnectionPolicyName.DEFAULT,
            new ServerConnectionPolicyInner().withConnectionType(ServerConnectionType.REDIRECT),
            com.azure.core.util.Context.NONE);
    }
}
