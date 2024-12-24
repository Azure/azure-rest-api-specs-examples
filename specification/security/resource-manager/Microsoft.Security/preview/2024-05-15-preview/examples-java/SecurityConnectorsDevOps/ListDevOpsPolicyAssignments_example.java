
/**
 * Samples for DevOpsPolicyAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/ListDevOpsPolicyAssignments_example.json
     */
    /**
     * Sample code: List_DevOpsPolicyAssignments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listDevOpsPolicyAssignments(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsPolicyAssignments().list("myRg", "mySecurityConnectorName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
