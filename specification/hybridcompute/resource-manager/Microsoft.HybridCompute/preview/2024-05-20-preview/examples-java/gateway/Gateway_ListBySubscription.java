
/**
 * Samples for Gateways List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/gateway/
     * Gateway_ListBySubscription.json
     */
    /**
     * Sample code: List Gateways by Subscription.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        listGatewaysBySubscription(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.gateways().list(com.azure.core.util.Context.NONE);
    }
}
