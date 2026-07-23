
/**
 * Samples for MongoMIResources ListMongoMIRoleAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleAssignmentList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoMIRoleAssignmentList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().listMongoMIRoleAssignments("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
