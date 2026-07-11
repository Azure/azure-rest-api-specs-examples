
/**
 * Samples for VolumeGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-15-preview/VolumeGroups_Get_Custom.json
     */
    /**
     * Sample code: VolumeGroups_Get_Custom.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsGetCustom(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().getWithResponse("myRG", "account1", "group1", com.azure.core.util.Context.NONE);
    }
}
