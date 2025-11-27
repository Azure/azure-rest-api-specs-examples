
/**
 * Samples for BackupsLongTermRetention Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * BackupsLongTermRetentionGet.json
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
