import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetExtensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtensions_Get_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetExtensionsGetMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetExtensions()
            .getWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", "aaaaaaa", Context.NONE);
    }
}
