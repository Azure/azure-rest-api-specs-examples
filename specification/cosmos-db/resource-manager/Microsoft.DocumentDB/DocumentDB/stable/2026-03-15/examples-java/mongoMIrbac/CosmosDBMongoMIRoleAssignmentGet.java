
/**
 * Samples for MongoMIResources GetMongoMIRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleAssignmentGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoMIRoleAssignmentGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().getMongoMIRoleAssignmentWithResponse("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
