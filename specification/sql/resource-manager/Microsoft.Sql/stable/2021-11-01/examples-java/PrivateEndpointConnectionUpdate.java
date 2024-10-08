
import com.azure.resourcemanager.sql.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.sql.models.PrivateLinkServiceConnectionStateProperty;
import com.azure.resourcemanager.sql.models.PrivateLinkServiceConnectionStateStatus;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        approveOrRejectAPrivateEndpointConnectionWithAGivenName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getPrivateEndpointConnections().createOrUpdate("Default",
            "test-svr", "private-endpoint-connection-name",
            new PrivateEndpointConnectionInner()
                .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionStateProperty()
                    .withStatus(PrivateLinkServiceConnectionStateStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com")),
            com.azure.core.util.Context.NONE);
    }
}
