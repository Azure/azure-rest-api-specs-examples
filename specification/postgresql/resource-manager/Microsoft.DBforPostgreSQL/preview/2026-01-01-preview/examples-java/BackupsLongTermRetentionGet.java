
/**
 * Samples for BackupsLongTermRetention Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/BackupsLongTermRetentionGet.json
     */
    /**
     * Sample code: Get the results of a long retention backup operation for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getTheResultsOfALongRetentionBackupOperationForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsLongTermRetentions().getWithResponse("exampleresourcegroup", "exampleserver", "exampleltrbackup",
            com.azure.core.util.Context.NONE);
    }
}
