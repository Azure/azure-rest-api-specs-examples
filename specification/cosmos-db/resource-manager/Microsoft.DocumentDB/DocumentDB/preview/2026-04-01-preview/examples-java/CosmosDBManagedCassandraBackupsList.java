
/**
 * Samples for CassandraClusters ListBackups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBManagedCassandraBackupsList.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraBackupsList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraBackupsList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().listBackups("cassandra-prod-rg", "cassandra-prod",
            com.azure.core.util.Context.NONE);
    }
}
