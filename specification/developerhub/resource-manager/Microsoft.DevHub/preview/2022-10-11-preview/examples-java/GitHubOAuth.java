import com.azure.resourcemanager.devhub.models.GitHubOAuthCallRequest;

/** Samples for ResourceProvider GitHubOAuth. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuth.json
     */
    /**
     * Sample code: GitHub OAuth.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void gitHubOAuth(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager
            .resourceProviders()
            .gitHubOAuthWithResponse(
                "eastus2euap",
                new GitHubOAuthCallRequest().withRedirectUrl("https://ms.portal.azure.com/aks"),
                com.azure.core.util.Context.NONE);
    }
}
