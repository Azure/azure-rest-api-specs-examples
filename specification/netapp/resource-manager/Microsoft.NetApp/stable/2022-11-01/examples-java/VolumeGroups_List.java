/** Samples for VolumeGroups ListByNetAppAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/VolumeGroups_List.json
     */
    /**
     * Sample code: VolumeGroups_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().listByNetAppAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
