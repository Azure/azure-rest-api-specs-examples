import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs PowerOff. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_PowerOff_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsPowerOffMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .powerOff("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaa", null, Context.NONE);
    }
}
