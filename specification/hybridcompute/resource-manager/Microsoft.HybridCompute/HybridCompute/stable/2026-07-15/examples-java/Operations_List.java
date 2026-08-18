
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/Operations_List.json
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
