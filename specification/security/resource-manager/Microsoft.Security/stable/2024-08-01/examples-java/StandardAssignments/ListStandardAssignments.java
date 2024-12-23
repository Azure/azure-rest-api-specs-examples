
/**
 * Samples for StandardAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/StandardAssignments/
     * ListStandardAssignments.json
     */
    /**
     * Sample code: List standard assignments.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listStandardAssignments(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.standardAssignments().list(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector",
            com.azure.core.util.Context.NONE);
    }
}
