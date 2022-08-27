import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetRollingUpgrades Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrades_Cancel_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrades_Cancel_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetRollingUpgradesCancelMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetRollingUpgrades()
            .cancel("rgcompute", "aaaaaa", Context.NONE);
    }
}
