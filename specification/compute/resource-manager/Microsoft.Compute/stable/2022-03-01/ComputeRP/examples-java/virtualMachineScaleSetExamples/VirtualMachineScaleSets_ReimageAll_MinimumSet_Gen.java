import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets ReimageAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ReimageAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ReimageAll_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsReimageAllMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .reimageAll("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", null, Context.NONE);
    }
}
