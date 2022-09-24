import com.azure.core.util.Context;

/** Samples for MongoDBResources DeleteMongoUserDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBMongoDBUserDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBUserDefinitionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .deleteMongoUserDefinition("myMongoUserDefinitionId", "myResourceGroupName", "myAccountName", Context.NONE);
    }
}
