import com.azure.resourcemanager.postgresqlflexibleserver.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.postgresqlflexibleserver.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.postgresqlflexibleserver.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnectionOperation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void approveOrRejectAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .privateEndpointConnectionOperations()
            .update(
                "Default",
                "test-svr",
                "private-endpoint-connection-name.1fa229cd-bf3f-47f0-8c49-afb36723997e",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("Approved by johndoe@contoso.com")),
                com.azure.core.util.Context.NONE);
    }
}
