
/**
 * Samples for MongoDBResources ListMongoUserDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBMongoDBUserDefinitionList.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBUserDefinitionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources()
            .listMongoUserDefinitions("myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
