
/**
 * Samples for MongoDBResources DeleteMongoRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBRoleDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBRoleDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().deleteMongoRoleDefinition("myMongoRoleDefinitionId",
            "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
