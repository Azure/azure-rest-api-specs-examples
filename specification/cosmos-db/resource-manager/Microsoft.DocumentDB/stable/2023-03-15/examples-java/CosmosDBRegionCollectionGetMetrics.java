/** Samples for CollectionRegion ListMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBRegionCollectionGetMetrics.json
     */
    /**
     * Sample code: CosmosDBRegionCollectionGetMetrics.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRegionCollectionGetMetrics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCollectionRegions()
            .listMetrics(
                "rg1",
                "ddb1",
                "North Europe",
                "databaseRid",
                "collectionRid",
                "$filter=(name.value eq 'Total Requests') and timeGrain eq duration'PT5M' and startTime eq"
                    + " '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z",
                com.azure.core.util.Context.NONE);
    }
}
