
/**
 * Samples for Gateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/gateway/Gateway_ListByResourceGroup.json
     */
    /**
     * Sample code: List Gateways by Resource Group.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        listGatewaysByResourceGroup(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.gateways().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
