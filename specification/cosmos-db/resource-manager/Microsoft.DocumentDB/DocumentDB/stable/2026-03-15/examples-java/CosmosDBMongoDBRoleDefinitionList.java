
/**
 * Samples for MongoDBResources ListMongoRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBMongoDBRoleDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBRoleDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().listMongoRoleDefinitions("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
