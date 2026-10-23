
/**
 * Samples for LogFiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/LogFiles_Get.json
     */
    /**
     * Sample code: Get captured log file metadata.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getCapturedLogFileMetadata(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.logFiles().getWithResponse("example-resource-group", "example-cluster",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
