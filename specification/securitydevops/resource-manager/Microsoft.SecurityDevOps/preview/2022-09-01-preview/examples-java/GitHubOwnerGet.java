import com.azure.core.util.Context;

/** Samples for GitHubOwner Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerGet.json
     */
    /**
     * Sample code: GitHubOwner_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubOwnerGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubOwners().getWithResponse("westusrg", "testconnector", "Azure", Context.NONE);
    }
}
