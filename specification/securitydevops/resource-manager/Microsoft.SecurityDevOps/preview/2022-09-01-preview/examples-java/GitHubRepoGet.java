import com.azure.core.util.Context;

/** Samples for GitHubRepo Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoGet.json
     */
    /**
     * Sample code: GitHubRepo_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubRepoGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubRepoes().getWithResponse("westusrg", "testconnector", "Azure", "39093389", Context.NONE);
    }
}
