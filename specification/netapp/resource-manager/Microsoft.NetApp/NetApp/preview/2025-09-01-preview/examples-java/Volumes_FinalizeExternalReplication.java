
/**
 * Samples for Volumes FinalizeExternalReplication.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_FinalizeExternalReplication.json
     */
    /**
     * Sample code: Volumes_FinalizeExternalReplication.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesFinalizeExternalReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().finalizeExternalReplication("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
