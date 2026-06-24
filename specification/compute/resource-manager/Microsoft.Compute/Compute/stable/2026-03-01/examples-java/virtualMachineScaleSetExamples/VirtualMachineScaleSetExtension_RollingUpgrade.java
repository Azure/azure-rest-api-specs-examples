
/**
 * Samples for VirtualMachineScaleSetRollingUpgrades StartExtensionUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_RollingUpgrade.json
     */
    /**
     * Sample code: Start an extension rolling upgrade.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void startAnExtensionRollingUpgrade(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetRollingUpgrades().startExtensionUpgrade("myResourceGroup",
            "{vmss-name}", com.azure.core.util.Context.NONE);
    }
}
