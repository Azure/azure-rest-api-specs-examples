
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades StartExtensionUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_RollingUpgrade.json
     */
    /**
     * Sample code: Start an extension rolling upgrade.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startAnExtensionRollingUpgrade(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetRollingUpgrades()
            .startExtensionUpgrade("myResourceGroup", "{vmss-name}", com.azure.core.util.Context.NONE);
    }
}
