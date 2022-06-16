import com.azure.core.util.Context;

/** Samples for VolumeGroups ListByNetAppAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/VolumeGroups_List.json
     */
    /**
     * Sample code: VolumeGroups_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().listByNetAppAccount("myRG", "account1", Context.NONE);
    }
}
