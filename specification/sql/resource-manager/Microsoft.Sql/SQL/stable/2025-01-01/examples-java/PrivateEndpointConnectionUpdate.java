
import com.azure.resourcemanager.sql.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.sql.models.PrivateLinkServiceConnectionStateProperty;
import com.azure.resourcemanager.sql.models.PrivateLinkServiceConnectionStateStatus;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void approveOrRejectAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().createOrUpdate("Default", "test-svr",
            "private-endpoint-connection-name",
            new PrivateEndpointConnectionInner()
                .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionStateProperty()
                    .withStatus(PrivateLinkServiceConnectionStateStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com")),
            com.azure.core.util.Context.NONE);
    }
}
