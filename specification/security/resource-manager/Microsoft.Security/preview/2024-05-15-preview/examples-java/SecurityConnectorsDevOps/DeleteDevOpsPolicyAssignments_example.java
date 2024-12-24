
/**
 * Samples for DevOpsPolicyAssignments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/DeleteDevOpsPolicyAssignments_example.json
     */
    /**
     * Sample code: Delete_DevOpsPolicyAssignments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteDevOpsPolicyAssignments(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsPolicyAssignments().delete("myRg", "mySecurityConnectorName",
            "5ec87f43-62d8-437b-8f46-4c8d4032cf6d", com.azure.core.util.Context.NONE);
    }
}
