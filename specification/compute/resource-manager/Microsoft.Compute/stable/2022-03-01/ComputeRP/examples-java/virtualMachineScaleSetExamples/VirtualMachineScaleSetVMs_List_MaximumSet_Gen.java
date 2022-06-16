import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsListMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .list(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaa",
                Context.NONE);
    }
}
