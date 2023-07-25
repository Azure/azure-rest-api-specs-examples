/** Samples for VolumeGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/VolumeGroups_Get.json
     */
    /**
     * Sample code: VolumeGroups_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().getWithResponse("myRG", "account1", "group1", com.azure.core.util.Context.NONE);
    }
}
