import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetRollingUpgrades StartOSUpgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrades_StartOSUpgrade_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrades_StartOSUpgrade_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetRollingUpgradesStartOSUpgradeMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetRollingUpgrades()
            .startOSUpgrade("rgcompute", "aaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
