import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtensions_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetExtensionsListMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetExtensions()
            .list("rgcompute", "aaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
