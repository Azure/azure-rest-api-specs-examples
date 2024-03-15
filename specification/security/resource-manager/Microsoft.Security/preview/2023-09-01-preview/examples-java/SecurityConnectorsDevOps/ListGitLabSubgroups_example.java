
/**
 * Samples for GitLabSubgroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/ListGitLabSubgroups_example.json
     */
    /**
     * Sample code: List_GitLabSubgroups.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listGitLabSubgroups(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitLabSubgroups().listWithResponse("myRg", "mySecurityConnectorName", "myGitLabGroup",
            com.azure.core.util.Context.NONE);
    }
}
