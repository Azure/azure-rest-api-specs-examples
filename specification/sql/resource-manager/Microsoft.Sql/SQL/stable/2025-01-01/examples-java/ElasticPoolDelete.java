
/**
 * Samples for ElasticPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ElasticPoolDelete.json
     */
    /**
     * Sample code: Delete an elastic pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAnElasticPool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getElasticPools().delete("sqlcrudtest-3129", "sqlcrudtest-228", "sqlcrudtest-3851",
            com.azure.core.util.Context.NONE);
    }
}
