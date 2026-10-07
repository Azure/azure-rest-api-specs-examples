
/**
 * Samples for CassandraResources DeleteCassandraView.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCassandraViewDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraViewDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraViewDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().deleteCassandraView("rg1", "ddb1", "keyspacename", "viewname",
            com.azure.core.util.Context.NONE);
    }
}
