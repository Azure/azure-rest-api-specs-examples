
/**
 * Samples for MongoMIResources ListMongoMIRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoMIRoleDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().listMongoMIRoleDefinitions("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
