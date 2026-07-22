
/**
 * Samples for MongoDBResources DeleteMongoUserDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBUserDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBUserDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().deleteMongoUserDefinition("myMongoUserDefinitionId",
            "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
