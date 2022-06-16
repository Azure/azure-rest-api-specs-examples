import com.azure.core.util.Context;

/** Samples for Volumes ReInitializeReplication. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Volumes_ReInitializeReplication.json
     */
    /**
     * Sample code: Volumes_ReInitializeReplication.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesReInitializeReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().reInitializeReplication("myRG", "account1", "pool1", "volume1", Context.NONE);
    }
}
