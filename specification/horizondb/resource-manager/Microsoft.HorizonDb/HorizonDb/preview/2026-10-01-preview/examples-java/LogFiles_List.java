
/**
 * Samples for LogFiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/LogFiles_List.json
     */
    /**
     * Sample code: List captured log files.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listCapturedLogFiles(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.logFiles().list("example-resource-group", "example-cluster", com.azure.core.util.Context.NONE);
    }
}
