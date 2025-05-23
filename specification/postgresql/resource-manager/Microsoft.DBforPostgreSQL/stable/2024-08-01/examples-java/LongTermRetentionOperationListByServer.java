
/**
 * Samples for LtrBackupOperations ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/
     * LongTermRetentionOperationListByServer.json
     */
    /**
     * Sample code: Sample List of Long Tern Retention Operations by Flexible Server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void sampleListOfLongTernRetentionOperationsByFlexibleServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.ltrBackupOperations().listByServer("rgLongTermRetention", "pgsqlltrtestserver",
            com.azure.core.util.Context.NONE);
    }
}
