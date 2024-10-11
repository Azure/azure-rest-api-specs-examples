
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/privateEndpoint/
     * PrivateEndpointConnection_Get.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        getsPrivateEndpointConnection(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.privateEndpointConnections().getWithResponse("myResourceGroup", "myPrivateLinkScope",
            "private-endpoint-connection-name", com.azure.core.util.Context.NONE);
    }
}
