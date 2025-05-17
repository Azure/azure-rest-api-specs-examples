
/**
 * Samples for Capabilities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Capabilities_Get.json
     */
    /**
     * Sample code: Get a Capability that extends a virtual machine Target resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getACapabilityThatExtendsAVirtualMachineTargetResource(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.capabilities().getWithResponse("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM",
            "Microsoft-VirtualMachine", "Shutdown-1.0", com.azure.core.util.Context.NONE);
    }
}
