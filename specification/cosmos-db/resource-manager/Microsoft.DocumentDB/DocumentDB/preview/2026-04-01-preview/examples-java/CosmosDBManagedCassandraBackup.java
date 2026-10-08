
/**
 * Samples for CassandraClusters GetBackup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBManagedCassandraBackup.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraBackup.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraBackup(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().getBackupWithResponse("cassandra-prod-rg", "cassandra-prod",
            "1611250348", com.azure.core.util.Context.NONE);
    }
}
