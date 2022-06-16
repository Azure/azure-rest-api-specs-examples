import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetRollingUpgrades StartExtensionUpgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VMScaleSetExtensionRollingUpgrade.json
     */
    /**
     * Sample code: Start an extension rolling upgrade.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startAnExtensionRollingUpgrade(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetRollingUpgrades()
            .startExtensionUpgrade("myResourceGroup", "{vmss-name}", Context.NONE);
    }
}
