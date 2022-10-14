import com.azure.core.util.Context;
import com.azure.resourcemanager.securitydevops.models.GitHubRepo;

/** Samples for GitHubRepo Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoUpdate.json
     */
    /**
     * Sample code: GitHubRepo_Update.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubRepoUpdate(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        GitHubRepo resource =
            manager
                .gitHubRepoes()
                .getWithResponse("westusrg", "testconnector", "Azure", "azure-rest-api-specs", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
