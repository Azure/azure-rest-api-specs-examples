
/**
 * Samples for LtrBackupOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/
     * LongTermRetentionOperationGet.json
     */
    /**
     * Sample code: Sample.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void sample(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.ltrBackupOperations().getWithResponse("rgLongTermRetention", "pgsqlltrtestserver", "backup1",
            com.azure.core.util.Context.NONE);
    }
}
