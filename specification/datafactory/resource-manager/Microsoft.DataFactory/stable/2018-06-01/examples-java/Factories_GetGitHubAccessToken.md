Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.9/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.datafactory.models.GitHubAccessTokenRequest;

/** Samples for Factories GetGitHubAccessToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetGitHubAccessToken.json
     */
    /**
     * Sample code: Factories_GetGitHubAccessToken.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesGetGitHubAccessToken(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .factories()
            .getGitHubAccessTokenWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                new GitHubAccessTokenRequest()
                    .withGitHubAccessCode("some")
                    .withGitHubClientId("some")
                    .withGitHubAccessTokenBaseUrl("some"),
                Context.NONE);
    }
}
```
