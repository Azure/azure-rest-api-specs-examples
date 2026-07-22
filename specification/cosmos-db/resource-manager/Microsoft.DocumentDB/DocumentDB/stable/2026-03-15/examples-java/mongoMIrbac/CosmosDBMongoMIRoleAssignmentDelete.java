
/**
 * Samples for MongoMIResources DeleteMongoMIRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleAssignmentDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoMIRoleAssignmentDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().deleteMongoMIRoleAssignment("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
