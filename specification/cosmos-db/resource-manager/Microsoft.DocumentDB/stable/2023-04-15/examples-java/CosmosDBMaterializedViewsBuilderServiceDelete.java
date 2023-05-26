/** Samples for Service Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBMaterializedViewsBuilderServiceDelete.json
     */
    /**
     * Sample code: MaterializedViewsBuilderServiceDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void materializedViewsBuilderServiceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getServices()
            .delete("rg1", "ddb1", "MaterializedViewsBuilder", com.azure.core.util.Context.NONE);
    }
}
