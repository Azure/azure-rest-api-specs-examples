/** Samples for VirtualMachineScaleSets GetOSUpgradeHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetOSUpgradeHistory_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_GetOSUpgradeHistory_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetGetOSUpgradeHistoryMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .getOSUpgradeHistory("rgcompute", "aaaaaa", com.azure.core.util.Context.NONE);
    }
}
