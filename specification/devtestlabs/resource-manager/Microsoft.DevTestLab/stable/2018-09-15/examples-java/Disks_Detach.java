import com.azure.resourcemanager.devtestlabs.models.DetachDiskProperties;

/** Samples for Disks Detach. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_Detach.json
     */
    /**
     * Sample code: Disks_Detach.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void disksDetach(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .disks()
            .detach(
                "resourceGroupName",
                "{labName}",
                "{userId}",
                "{diskName}",
                new DetachDiskProperties()
                    .withLeasedByLabVmId(
                        "/subscriptions/{subscriptionId}/resourcegroups/myResourceGroup/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/{vmName}"),
                com.azure.core.util.Context.NONE);
    }
}
