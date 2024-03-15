
/**
 * Samples for Capabilities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/DeleteCapability.json
     */
    /**
     * Sample code: Delete a Capability that extends a virtual machine Target resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteACapabilityThatExtendsAVirtualMachineTargetResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.capabilities().deleteWithResponse("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM",
            "Microsoft-VirtualMachine", "Shutdown-1.0", com.azure.core.util.Context.NONE);
    }
}
