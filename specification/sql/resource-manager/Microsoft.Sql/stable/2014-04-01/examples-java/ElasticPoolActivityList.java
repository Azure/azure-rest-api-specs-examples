import com.azure.core.util.Context;

/** Samples for ElasticPoolActivities ListByElasticPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ElasticPoolActivityList.json
     */
    /**
     * Sample code: List Elastic pool activity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listElasticPoolActivity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getElasticPoolActivities()
            .listByElasticPool("sqlcrudtest-4291", "sqlcrudtest-6574", "8749", Context.NONE);
    }
}
