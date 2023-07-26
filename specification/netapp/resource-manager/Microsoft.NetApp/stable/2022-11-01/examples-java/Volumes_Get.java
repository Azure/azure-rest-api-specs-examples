/** Samples for Volumes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Volumes_Get.json
     */
    /**
     * Sample code: Volumes_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().getWithResponse("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
