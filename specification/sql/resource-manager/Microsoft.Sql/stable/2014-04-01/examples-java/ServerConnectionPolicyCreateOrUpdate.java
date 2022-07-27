import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerConnectionPolicyInner;
import com.azure.resourcemanager.sql.models.ConnectionPolicyName;
import com.azure.resourcemanager.sql.models.ServerConnectionType;

/** Samples for ServerConnectionPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerConnectionPolicyCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a server's secure connection policy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAServerSSecureConnectionPolicy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerConnectionPolicies()
            .createOrUpdateWithResponse(
                "test-1234",
                "test-5678",
                ConnectionPolicyName.DEFAULT,
                new ServerConnectionPolicyInner().withConnectionType(ServerConnectionType.PROXY),
                Context.NONE);
    }
}
