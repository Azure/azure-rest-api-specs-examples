
/**
 * Samples for MongoDBResources GetMongoRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/
     * CosmosDBMongoDBRoleDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBMongoRoleDefinitionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoRoleDefinitionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().getMongoRoleDefinitionWithResponse(
            "myMongoRoleDefinitionId", "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
