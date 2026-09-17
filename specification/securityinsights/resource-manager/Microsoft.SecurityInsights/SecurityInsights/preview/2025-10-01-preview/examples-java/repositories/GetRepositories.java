
import com.azure.resourcemanager.securityinsights.models.RepositoryAccessKind;
import com.azure.resourcemanager.securityinsights.models.RepositoryAccessProperties;

/**
 * Samples for SourceControlOperation ListRepositories.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/repositories/GetRepositories.json
     */
    /**
     * Sample code: Get repository list.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getRepositoryList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControlOperations().listRepositories("myRg", "myWorkspace",
            new RepositoryAccessProperties().withKind(RepositoryAccessKind.OAUTH).withCode("fakeTokenPlaceholder")
                .withState("state").withClientId("54b3c2c0-1f48-4a1c-af9f-6399c3240b73"),
            com.azure.core.util.Context.NONE);
    }
}
