
/**
 * Samples for MachineExtensions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * extension/Extension_List.json
     */
    /**
     * Sample code: GET all Machine Extensions - List.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        gETAllMachineExtensionsList(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineExtensions().list("myResourceGroup", "myMachine", null, com.azure.core.util.Context.NONE);
    }
}
