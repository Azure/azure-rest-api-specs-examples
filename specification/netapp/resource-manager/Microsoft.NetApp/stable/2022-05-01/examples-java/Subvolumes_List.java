import com.azure.core.util.Context;

/** Samples for Subvolumes ListByVolume. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/Subvolumes_List.json
     */
    /**
     * Sample code: Subvolumes_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().listByVolume("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
