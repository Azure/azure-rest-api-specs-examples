
/**
 * Samples for VirtualMachineScaleSetVMs PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_PowerOff_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMPowerOffMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().powerOff("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
