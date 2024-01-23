
import com.azure.resourcemanager.datafactory.models.GitHubAccessTokenRequest;

/**
 * Samples for Factories GetGitHubAccessToken.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Factories_GetGitHubAccessToken.json
     */
    /**
     * Sample code: Factories_GetGitHubAccessToken.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesGetGitHubAccessToken(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().getGitHubAccessTokenWithResponse("exampleResourceGroup", "exampleFactoryName",
            new GitHubAccessTokenRequest().withGitHubAccessCode("fakeTokenPlaceholder").withGitHubClientId("some")
                .withGitHubAccessTokenBaseUrl("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
