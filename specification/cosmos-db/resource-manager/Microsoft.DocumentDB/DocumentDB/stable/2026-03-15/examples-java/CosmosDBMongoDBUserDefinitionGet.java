
/**
 * Samples for MongoDBResources GetMongoUserDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBUserDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBUserDefinitionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().getMongoUserDefinitionWithResponse("myMongoUserDefinitionId",
            "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
