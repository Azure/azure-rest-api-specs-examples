
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/
     * Operations_List.json
     */
    /**
     * Sample code: List Hybrid Compute Provider Operations.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        listHybridComputeProviderOperations(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
