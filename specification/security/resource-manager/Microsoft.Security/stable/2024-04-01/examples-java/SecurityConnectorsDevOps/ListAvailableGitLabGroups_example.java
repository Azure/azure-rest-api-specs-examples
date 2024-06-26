
/**
 * Samples for GitLabGroups ListAvailable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListAvailableGitLabGroups_example.json
     */
    /**
     * Sample code: ListAvailable_GitLabGroups.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAvailableGitLabGroups(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitLabGroups().listAvailableWithResponse("myRg", "mySecurityConnectorName",
            com.azure.core.util.Context.NONE);
    }
}
