import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetRollingUpgrades GetLatest. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetRollingUpgrades_GetLatest_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetRollingUpgrades_GetLatest_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetRollingUpgradesGetLatestMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetRollingUpgrades()
            .getLatestWithResponse("rgcompute", "aaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
