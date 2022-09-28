import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.RepoType;

/** Samples for SourceControl ListRepositories. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/repositories/GetRepositories.json
     */
    /**
     * Sample code: Get repository list.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getRepositoryList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControls().listRepositories("myRg", "myWorkspace", RepoType.GITHUB, Context.NONE);
    }
}
