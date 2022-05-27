Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BuildService GetSupportedStack. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetSupportedStack.json
     */
    /**
     * Sample code: BuildService_GetSupportedStack.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetSupportedStack(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServices()
            .getSupportedStackWithResponse(
                "myResourceGroup", "myservice", "default", "io.buildpacks.stacks.bionic-base", Context.NONE);
    }
}
```
