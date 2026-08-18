
/**
 * Samples for MachineExtensions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/extension/Extension_List.json
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
