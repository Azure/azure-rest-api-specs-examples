
/**
 * Samples for Volumes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_ExtraLargeVolumes_List.json
     */
    /**
     * Sample code: Volumes_ExtralargeVolumeList.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesExtralargeVolumeList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().list("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
