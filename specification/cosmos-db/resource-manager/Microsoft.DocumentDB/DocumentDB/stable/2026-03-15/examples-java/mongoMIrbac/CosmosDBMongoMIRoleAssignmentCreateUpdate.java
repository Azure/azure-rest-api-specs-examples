
import com.azure.resourcemanager.cosmos.fluent.models.MongoMIRoleAssignmentResourceInner;

/**
 * Samples for MongoMIResources CreateUpdateMongoMIRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/mongoMIrbac/CosmosDBMongoMIRoleAssignmentCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoMIRoleAssignmentCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoMIRoleAssignmentCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoMIResources().createUpdateMongoMIRoleAssignment("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId",
            new MongoMIRoleAssignmentResourceInner().withRoleDefinitionId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongoMIRoleDefinitions/myRoleDefinitionId")
                .withScope(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases")
                .withPrincipalId("myPrincipalId"),
            com.azure.core.util.Context.NONE);
    }
}
