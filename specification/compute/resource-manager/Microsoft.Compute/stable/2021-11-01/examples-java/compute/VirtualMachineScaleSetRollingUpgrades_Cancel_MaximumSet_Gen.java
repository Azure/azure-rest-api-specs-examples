import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetRollingUpgrades Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetRollingUpgrades_Cancel_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrades_Cancel_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetRollingUpgradesCancelMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetRollingUpgrades()
            .cancel("rgcompute", "aaaaa", Context.NONE);
    }
}
