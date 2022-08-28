import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_Start_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsStartMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .start("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
