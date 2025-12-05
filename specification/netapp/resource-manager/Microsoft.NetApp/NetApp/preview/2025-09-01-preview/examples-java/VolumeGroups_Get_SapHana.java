
/**
 * Samples for VolumeGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/VolumeGroups_Get_SapHana.json
     */
    /**
     * Sample code: VolumeGroups_Get_SapHana.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsGetSapHana(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().getWithResponse("myRG", "account1", "group1", com.azure.core.util.Context.NONE);
    }
}
