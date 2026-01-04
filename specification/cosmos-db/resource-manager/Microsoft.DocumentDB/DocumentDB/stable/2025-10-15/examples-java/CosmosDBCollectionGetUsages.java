
/**
 * Samples for Collection ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBCollectionGetUsages.json
     */
    /**
     * Sample code: CosmosDBCollectionGetUsages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCollectionGetUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCollections().listUsages("rg1", "ddb1", "databaseRid",
            "collectionRid", "$filter=name.value eq 'Storage'", com.azure.core.util.Context.NONE);
    }
}
