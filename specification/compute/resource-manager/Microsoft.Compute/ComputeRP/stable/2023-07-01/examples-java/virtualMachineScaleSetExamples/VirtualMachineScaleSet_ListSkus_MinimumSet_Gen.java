/** Samples for VirtualMachineScaleSets ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ListSkus_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetListSkusMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .listSkus("rgcompute", "aaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
