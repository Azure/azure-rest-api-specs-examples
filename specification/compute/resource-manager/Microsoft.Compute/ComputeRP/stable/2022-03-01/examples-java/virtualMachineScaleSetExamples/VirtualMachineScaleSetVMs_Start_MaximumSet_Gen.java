import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_Start_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_Start_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsStartMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .start("rgcompute", "aaaaaaaaaaaaaa", "aaaaaaaaaaaaa", Context.NONE);
    }
}
