
/**
 * Samples for CassandraResources ListCassandraKeyspaces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraKeyspaceList.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraKeyspaceList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().listCassandraKeyspaces("rgName", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
