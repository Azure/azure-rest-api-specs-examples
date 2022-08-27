import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_Reimage_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Reimage_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsReimageMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .reimage("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", null, Context.NONE);
    }
}
