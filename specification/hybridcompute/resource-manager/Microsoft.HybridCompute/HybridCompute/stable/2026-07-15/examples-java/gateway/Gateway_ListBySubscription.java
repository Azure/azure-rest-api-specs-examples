
/**
 * Samples for Gateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/gateway/Gateway_ListBySubscription.json
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
