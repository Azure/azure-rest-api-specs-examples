import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Start_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsStartMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .start("rgcompute", "aaaaaaaaaaaaaaaaaaa", null, Context.NONE);
    }
}
