/** Samples for VirtualMachineScaleSets Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Start_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetStartMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .start("rgcompute", "aaaaaaaaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
