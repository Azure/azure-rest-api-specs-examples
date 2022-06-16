import com.azure.core.util.Context;

/** Samples for DatabaseAccounts ListMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBDatabaseAccountGetMetrics.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountGetMetrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountGetMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabaseAccounts()
            .listMetrics(
                "rg1",
                "ddb1",
                "$filter=(name.value eq 'Total Requests') and timeGrain eq duration'PT5M' and startTime eq"
                    + " '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z",
                Context.NONE);
    }
}
