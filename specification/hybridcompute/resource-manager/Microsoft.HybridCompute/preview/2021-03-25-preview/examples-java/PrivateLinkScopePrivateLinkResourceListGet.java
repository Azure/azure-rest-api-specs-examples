/** Samples for PrivateLinkResources ListByPrivateLinkScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-03-25-preview/examples/PrivateLinkScopePrivateLinkResourceListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getsPrivateEndpointConnection(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager
            .privateLinkResources()
            .listByPrivateLinkScope("myResourceGroup", "myPrivateLinkScope", com.azure.core.util.Context.NONE);
    }
}
