import com.azure.core.util.Context;
import com.azure.resourcemanager.securitydevops.models.GitHubOwner;

/** Samples for GitHubOwner Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerUpdate.json
     */
    /**
     * Sample code: GitHubOwner_Update.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubOwnerUpdate(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        GitHubOwner resource =
            manager.gitHubOwners().getWithResponse("westusrg", "testconnector", "Azure", Context.NONE).getValue();
        resource.update().apply();
    }
}
