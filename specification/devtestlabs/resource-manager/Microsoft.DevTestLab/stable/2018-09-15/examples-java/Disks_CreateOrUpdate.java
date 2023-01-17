import com.azure.resourcemanager.devtestlabs.models.StorageType;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_CreateOrUpdate.json
     */
    /**
     * Sample code: Disks_CreateOrUpdate.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void disksCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .disks()
            .define("{diskName}")
            .withRegion((String) null)
            .withExistingUser("resourceGroupName", "{labName}", "{userId}")
            .withDiskType(StorageType.STANDARD)
            .withDiskSizeGiB(1023)
            .withLeasedByLabVmId(
                "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualmachines/vmName")
            .create();
    }
}
