
/**
 * Samples for GitLabGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListGitLabGroups_example.json
     */
    /**
     * Sample code: List_GitLabGroups.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listGitLabGroups(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitLabGroups().list("myRg", "mySecurityConnectorName", com.azure.core.util.Context.NONE);
    }
}
