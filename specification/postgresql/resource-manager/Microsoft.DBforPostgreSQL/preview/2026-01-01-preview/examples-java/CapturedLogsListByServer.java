
/**
 * Samples for CapturedLogs ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/CapturedLogsListByServer.json
     */
    /**
     * Sample code: List all captured logs for download in a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAllCapturedLogsForDownloadInAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.capturedLogs().listByServer("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
