import com.azure.core.util.Context;

/** Samples for Volumes DeleteReplication. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/Volumes_DeleteReplication.json
     */
    /**
     * Sample code: Volumes_DeleteReplication.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesDeleteReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().deleteReplication("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
