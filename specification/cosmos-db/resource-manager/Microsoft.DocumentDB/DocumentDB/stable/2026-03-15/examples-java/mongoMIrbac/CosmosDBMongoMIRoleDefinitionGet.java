
/**
 * Samples for MongoMIResources GetMongoMIRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleDefinitionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoMIRoleDefinitionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().getMongoMIRoleDefinitionWithResponse("myResourceGroupName",
            "myAccountName", "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
