/** Samples for PercentileTarget ListMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBPercentileTargetGetMetrics.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountRegionGetMetrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountRegionGetMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getPercentileTargets()
            .listMetrics(
                "rg1",
                "ddb1",
                "East US",
                "$filter=(name.value eq 'Probabilistic Bounded Staleness') and timeGrain eq duration'PT5M' and"
                    + " startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z",
                com.azure.core.util.Context.NONE);
    }
}
