
/**
 * Samples for Gateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/gateway/
     * Gateway_Delete.json
     */
    /**
     * Sample code: Delete a Gateway.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteAGateway(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.gateways().delete("myResourceGroup", "{gatewayName}", com.azure.core.util.Context.NONE);
    }
}
