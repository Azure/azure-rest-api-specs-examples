
/**
 * Samples for Volumes FinalizeRelocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_FinalizeRelocation.json
     */
    /**
     * Sample code: Volumes_FinalizeRelocation.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesFinalizeRelocation(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().finalizeRelocation("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
