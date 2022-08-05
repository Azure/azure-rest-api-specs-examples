import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets PowerOff. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_PowerOff_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsPowerOffMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .powerOff("rgcompute", "a", null, null, Context.NONE);
    }
}
