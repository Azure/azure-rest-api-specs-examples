
/**
 * Samples for VolumeGroups ListByNetAppAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * VolumeGroups_List_SapHana.json
     */
    /**
     * Sample code: VolumeGroups_List_SapHana.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsListSapHana(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().listByNetAppAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
