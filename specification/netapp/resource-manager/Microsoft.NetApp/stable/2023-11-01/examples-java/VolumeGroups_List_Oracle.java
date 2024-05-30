
/**
 * Samples for VolumeGroups ListByNetAppAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/VolumeGroups_List_Oracle.json
     */
    /**
     * Sample code: VolumeGroups_List_Oracle.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsListOracle(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().listByNetAppAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
