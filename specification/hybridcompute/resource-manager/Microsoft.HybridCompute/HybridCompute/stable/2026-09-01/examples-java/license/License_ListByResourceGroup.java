
/**
 * Samples for Licenses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/license/License_ListByResourceGroup.json
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
