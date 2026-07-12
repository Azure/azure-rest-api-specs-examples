
/**
 * Samples for VolumeGroups ListByNetAppAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-15-preview/VolumeGroups_List_Custom.json
     */
    /**
     * Sample code: VolumeGroups_List_Custom.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsListCustom(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().listByNetAppAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
