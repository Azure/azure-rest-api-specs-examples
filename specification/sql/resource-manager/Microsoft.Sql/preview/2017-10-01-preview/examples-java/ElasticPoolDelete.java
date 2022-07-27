import com.azure.core.util.Context;

/** Samples for ElasticPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ElasticPoolDelete.json
     */
    /**
     * Sample code: Delete an elastic pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnElasticPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getElasticPools()
            .delete("sqlcrudtest-3129", "sqlcrudtest-228", "sqlcrudtest-3851", Context.NONE);
    }
}
