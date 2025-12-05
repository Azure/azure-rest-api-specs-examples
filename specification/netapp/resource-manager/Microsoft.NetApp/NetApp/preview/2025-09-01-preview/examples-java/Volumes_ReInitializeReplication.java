
/**
 * Samples for Volumes ReInitializeReplication.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_ReInitializeReplication.json
     */
    /**
     * Sample code: Volumes_ReInitializeReplication.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesReInitializeReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().reInitializeReplication("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
