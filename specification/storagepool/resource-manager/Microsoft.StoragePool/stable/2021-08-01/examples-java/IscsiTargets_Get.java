import com.azure.core.util.Context;

/** Samples for IscsiTargets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Get.json
     */
    /**
     * Sample code: Get iSCSI Target.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void getISCSITarget(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.iscsiTargets().getWithResponse("myResourceGroup", "myDiskPool", "myIscsiTarget", Context.NONE);
    }
}
