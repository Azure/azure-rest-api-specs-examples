import com.azure.core.util.Context;

/** Samples for ResourceProvider ListGitHubOAuth. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-04-01-preview/examples/GitHubOAuth_List.json
     */
    /**
     * Sample code: List GitHub OAuth.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void listGitHubOAuth(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager.resourceProviders().listGitHubOAuthWithResponse("eastus2euap", Context.NONE);
    }
}
