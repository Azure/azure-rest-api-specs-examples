/** Samples for GitHubRepo CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoCreateOrUpdate.json
     */
    /**
     * Sample code: GitHubRepo_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubRepoCreateOrUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager
            .gitHubRepoes()
            .define("azure-rest-api-specs")
            .withExistingOwner("westusrg", "testconnector", "Azure")
            .create();
    }
}
