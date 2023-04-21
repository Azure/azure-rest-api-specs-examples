/** Samples for Service Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBMaterializedViewsBuilderServiceGet.json
     */
    /**
     * Sample code: MaterializedViewsBuilderServiceGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void materializedViewsBuilderServiceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getServices()
            .getWithResponse("rg1", "ddb1", "MaterializedViewsBuilder", com.azure.core.util.Context.NONE);
    }
}
