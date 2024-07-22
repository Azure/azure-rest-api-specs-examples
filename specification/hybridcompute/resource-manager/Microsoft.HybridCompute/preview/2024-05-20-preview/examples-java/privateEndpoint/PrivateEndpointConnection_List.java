
/**
 * Samples for PrivateEndpointConnections ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/
     * privateEndpoint/PrivateEndpointConnection_List.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a private link scope.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getsListOfPrivateEndpointConnectionsOnAPrivateLinkScope(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateEndpointConnections().listByPrivateLinkScope("myResourceGroup", "myPrivateLinkScope",
            com.azure.core.util.Context.NONE);
    }
}
