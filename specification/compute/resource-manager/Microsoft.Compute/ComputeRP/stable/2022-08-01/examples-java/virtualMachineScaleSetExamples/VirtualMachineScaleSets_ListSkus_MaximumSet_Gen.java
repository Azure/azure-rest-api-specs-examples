import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ListSkus_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ListSkus_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsListSkusMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .listSkus("rgcompute", "aaaaaa", Context.NONE);
    }
}
