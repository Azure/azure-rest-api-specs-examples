
import com.azure.resourcemanager.devtestlabs.models.DetachDataDiskProperties;

/**
 * Samples for VirtualMachines DetachDataDisk.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachines_DetachDataDisk.json
     */
    /**
     * Sample code: VirtualMachines_DetachDataDisk.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesDetachDataDisk(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().detachDataDisk("resourceGroupName", "{labName}", "{virtualMachineName}",
            new DetachDataDiskProperties().withExistingLabDiskId(
                "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{virtualMachineName}"),
            com.azure.core.util.Context.NONE);
    }
}
