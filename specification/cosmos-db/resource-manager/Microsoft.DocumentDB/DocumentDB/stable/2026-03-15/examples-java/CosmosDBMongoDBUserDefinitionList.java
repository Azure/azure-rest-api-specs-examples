
/**
 * Samples for MongoDBResources ListMongoUserDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBUserDefinitionList.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBUserDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().listMongoUserDefinitions("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
