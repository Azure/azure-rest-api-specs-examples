Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.ConfigServerGitProperty;
import com.azure.resourcemanager.appplatform.models.ConfigServerSettings;
import java.util.Arrays;

/** Samples for ConfigServers Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigServers_Validate.json
     */
    /**
     * Sample code: ConfigServers_Validate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configServersValidate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigServers()
            .validate(
                "myResourceGroup",
                "myservice",
                new ConfigServerSettings()
                    .withGitProperty(
                        new ConfigServerGitProperty()
                            .withUri("https://github.com/fake-user/fake-repository.git")
                            .withLabel("master")
                            .withSearchPaths(Arrays.asList("/"))),
                Context.NONE);
    }
}
```
