
/**
 * Samples for Gateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/gateway/Gateway_Delete.json
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
