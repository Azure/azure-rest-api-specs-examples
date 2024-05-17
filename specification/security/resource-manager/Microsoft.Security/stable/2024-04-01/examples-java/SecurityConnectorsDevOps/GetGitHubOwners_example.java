
/**
 * Samples for GitHubOwners Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * GetGitHubOwners_example.json
     */
    /**
     * Sample code: Get_GitHubOwners.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getGitHubOwners(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitHubOwners().getWithResponse("myRg", "mySecurityConnectorName", "myGitHubOwner",
            com.azure.core.util.Context.NONE);
    }
}
