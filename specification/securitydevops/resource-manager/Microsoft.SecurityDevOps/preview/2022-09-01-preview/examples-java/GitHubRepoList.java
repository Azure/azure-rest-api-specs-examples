import com.azure.core.util.Context;

/** Samples for GitHubRepo List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoList.json
     */
    /**
     * Sample code: GitHubRepo_List.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubRepoList(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubRepoes().list("westusrg", "testconnector", "Azure", Context.NONE);
    }
}
