import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs PowerOff. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_PowerOff_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_PowerOff_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsPowerOffMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .powerOff("rgcompute", "aaaaaa", "aaaaaaaaa", true, Context.NONE);
    }
}
