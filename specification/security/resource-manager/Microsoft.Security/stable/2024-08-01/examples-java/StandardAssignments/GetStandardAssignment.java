
/**
 * Samples for StandardAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/StandardAssignments/
     * GetStandardAssignment.json
     */
    /**
     * Sample code: Retrieve a standard assignment.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void retrieveAStandardAssignment(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.standardAssignments().getWithResponse("providers/Microsoft.Management/managementGroups/contoso",
            "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", com.azure.core.util.Context.NONE);
    }
}
