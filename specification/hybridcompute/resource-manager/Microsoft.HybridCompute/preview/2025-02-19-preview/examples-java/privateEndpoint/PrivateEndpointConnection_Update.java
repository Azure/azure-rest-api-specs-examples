
import com.azure.resourcemanager.hybridcompute.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.hybridcompute.models.PrivateLinkServiceConnectionStateProperty;

/**
 * Samples for PrivateEndpointConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * privateEndpoint/PrivateEndpointConnection_Update.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void approveOrRejectAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateEndpointConnections().define("private-endpoint-connection-name")
            .withExistingPrivateLinkScope("myResourceGroup", "myPrivateLinkScope")
            .withProperties(new PrivateEndpointConnectionProperties()
                .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionStateProperty()
                    .withStatus("Approved").withDescription("Approved by johndoe@contoso.com")))
            .create();
    }
}
