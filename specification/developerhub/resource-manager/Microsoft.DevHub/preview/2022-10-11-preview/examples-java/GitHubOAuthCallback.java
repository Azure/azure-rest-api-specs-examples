/** Samples for ResourceProvider GitHubOAuthCallback. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/GitHubOAuthCallback.json
     */
    /**
     * Sample code: GitHub OAuth Callback.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void gitHubOAuthCallback(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager
            .resourceProviders()
            .gitHubOAuthCallbackWithResponse(
                "eastus2euap",
                "3584d83530557fdd1f46af8289938c8ef79f9dc5",
                "12345678-3456-7890-5678-012345678901",
                com.azure.core.util.Context.NONE);
    }
}
