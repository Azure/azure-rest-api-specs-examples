
/**
 * Samples for BackupsLongTermRetention ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/BackupsLongTermRetentionListByServer.json
     */
    /**
     * Sample code: List the results of the long term retention backup operations for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheResultsOfTheLongTermRetentionBackupOperationsForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsLongTermRetentions().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
