import com.azure.core.util.Context;

/** Samples for GitHubOwner List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerList.json
     */
    /**
     * Sample code: GitHubOwner_List.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubOwnerList(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubOwners().list("westusrg", "testconnector", Context.NONE);
    }
}
