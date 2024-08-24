
import com.azure.resourcemanager.search.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.search.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.search.models.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState;
import com.azure.resourcemanager.search.models.PrivateLinkServiceConnectionStatus;

/**
 * Samples for PrivateEndpointConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/UpdatePrivateEndpointConnection
     * .json
     */
    /**
     * Sample code: PrivateEndpointConnectionUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndpointConnectionUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getPrivateEndpointConnections().updateWithResponse("rg1",
            "mysearchservice", "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546",
            new PrivateEndpointConnectionInner()
                .withProperties(new PrivateEndpointConnectionProperties().withPrivateLinkServiceConnectionState(
                    new PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState()
                        .withStatus(PrivateLinkServiceConnectionStatus.REJECTED)
                        .withDescription("Rejected for some reason"))),
            null, com.azure.core.util.Context.NONE);
    }
}
