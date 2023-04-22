/** Samples for MongoDBResources ListMongoRoleDefinitions. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBMongoDBRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBMongoDBRoleDefinitionList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBRoleDefinitionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .listMongoRoleDefinitions("myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
