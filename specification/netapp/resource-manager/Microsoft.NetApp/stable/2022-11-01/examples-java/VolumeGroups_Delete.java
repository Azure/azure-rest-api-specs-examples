/** Samples for VolumeGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/VolumeGroups_Delete.json
     */
    /**
     * Sample code: VolumeGroups_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().delete("myRG", "account1", "group1", com.azure.core.util.Context.NONE);
    }
}
