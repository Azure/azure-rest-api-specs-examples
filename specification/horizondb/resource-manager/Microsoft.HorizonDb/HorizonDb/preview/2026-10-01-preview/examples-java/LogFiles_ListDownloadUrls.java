
/**
 * Samples for LogFiles ListDownloadUrls.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/LogFiles_ListDownloadUrls.json
     */
    /**
     * Sample code: List a signed URL for downloading a captured log file.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listASignedURLForDownloadingACapturedLogFile(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.logFiles().listDownloadUrlsWithResponse("example-resource-group", "example-cluster",
            "22222222-2222-2222-2222-222222222222", com.azure.core.util.Context.NONE);
    }
}
