
/**
 * Samples for Licenses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/
     * License_ListByResourceGroup.json
     */
    /**
     * Sample code: GET all Machine Extensions.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAllMachineExtensions(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenses().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
