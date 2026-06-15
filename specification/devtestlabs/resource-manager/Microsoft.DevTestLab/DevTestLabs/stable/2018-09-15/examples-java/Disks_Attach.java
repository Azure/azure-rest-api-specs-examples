
import com.azure.resourcemanager.devtestlabs.models.AttachDiskProperties;

/**
 * Samples for Disks Attach.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Disks_Attach.json
     */
    /**
     * Sample code: Disks_Attach.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void disksAttach(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.disks().attach("resourceGroupName", "{labName}", "{userId}", "{diskName}",
            new AttachDiskProperties().withLeasedByLabVmId(
                "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
            com.azure.core.util.Context.NONE);
    }
}
