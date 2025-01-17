
/**
 * Samples for GitHubOwners ListAvailable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/ListAvailableGitHubOwners_example.json
     */
    /**
     * Sample code: ListAvailable_GitHubOwners.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAvailableGitHubOwners(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitHubOwners().listAvailableWithResponse("myRg", "mySecurityConnectorName",
            com.azure.core.util.Context.NONE);
    }
}
