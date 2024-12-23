
/**
 * Samples for DevOpsPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/ListDevOpsPolicies_example.json
     */
    /**
     * Sample code: List_DevOpsPolicies.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listDevOpsPolicies(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsPolicies().list("myRg", "mySecurityConnectorName", null, com.azure.core.util.Context.NONE);
    }
}
