
/**
 * Samples for Targets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Targets_Get.json
     */
    /**
     * Sample code: Get a Target that extends a virtual machine resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getATargetThatExtendsAVirtualMachineResource(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.targets().getWithResponse("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM",
            "Microsoft-Agent", com.azure.core.util.Context.NONE);
    }
}
