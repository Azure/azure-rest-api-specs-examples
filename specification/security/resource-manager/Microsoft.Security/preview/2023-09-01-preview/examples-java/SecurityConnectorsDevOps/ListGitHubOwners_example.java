
/**
 * Samples for GitHubOwners List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/ListGitHubOwners_example.json
     */
    /**
     * Sample code: List_GitHubOwners.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listGitHubOwners(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitHubOwners().list("myRg", "mySecurityConnectorName", com.azure.core.util.Context.NONE);
    }
}
