
/**
 * Samples for MachineExtensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-16-preview/extension/Extension_Get.json
     */
    /**
     * Sample code: GET Machine Extension.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETMachineExtension(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machineExtensions().getWithResponse("myResourceGroup", "myMachine", "CustomScriptExtension",
            com.azure.core.util.Context.NONE);
    }
}
