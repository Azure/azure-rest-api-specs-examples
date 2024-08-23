
/**
 * Samples for VirtualMachines PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExamples/VirtualMachine_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_PowerOff_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinePowerOffMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().powerOff("rgcompute",
            "aaaaaaaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
