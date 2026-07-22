
/**
 * Samples for CassandraResources ListCassandraTables.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraTableList.json
     */
    /**
     * Sample code: CosmosDBCassandraTableList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraTableList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().listCassandraTables("rgName", "ddb1", "keyspaceName",
            com.azure.core.util.Context.NONE);
    }
}
