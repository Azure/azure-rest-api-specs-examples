
/**
 * Samples for IscsiTargets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Delete.
     * json
     */
    /**
     * Sample code: Delete iSCSI Target.
     * 
     * @param manager Entry point to StoragePoolManager.
     */
    public static void deleteISCSITarget(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.iscsiTargets().delete("myResourceGroup", "myDiskPool", "myIscsiTarget",
            com.azure.core.util.Context.NONE);
    }
}
