import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetExtensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtensions_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetExtensionsDeleteMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetExtensions()
            .delete("rgcompute", "aaaa", "aaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
