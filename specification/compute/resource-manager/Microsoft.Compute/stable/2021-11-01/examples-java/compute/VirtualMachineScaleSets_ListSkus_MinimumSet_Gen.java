import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ListSkus_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsListSkusMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .listSkus("rgcompute", "aaaaaaaaaaaaaaaa", Context.NONE);
    }
}
