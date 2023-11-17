/** Samples for Capabilities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetCapability.json
     */
    /**
     * Sample code: Get a Capability that extends a virtual machine Target resource.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void getACapabilityThatExtendsAVirtualMachineTargetResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .capabilities()
            .getWithResponse(
                "exampleRG",
                "Microsoft.Compute",
                "virtualMachines",
                "exampleVM",
                "Microsoft-VirtualMachine",
                "Shutdown-1.0",
                com.azure.core.util.Context.NONE);
    }
}
