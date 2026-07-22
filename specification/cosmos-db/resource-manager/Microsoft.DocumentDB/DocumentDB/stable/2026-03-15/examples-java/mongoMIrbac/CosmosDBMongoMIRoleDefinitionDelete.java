
/**
 * Samples for MongoMIResources DeleteMongoMIRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoMIRoleDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().deleteMongoMIRoleDefinition("myResourceGroupName",
            "myAccountName", "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
