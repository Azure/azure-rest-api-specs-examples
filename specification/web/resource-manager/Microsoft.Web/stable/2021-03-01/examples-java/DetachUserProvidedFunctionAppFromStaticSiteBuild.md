Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites DetachUserProvidedFunctionAppFromStaticSiteBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DetachUserProvidedFunctionAppFromStaticSiteBuild.json
     */
    /**
     * Sample code: Detach the user provided function app from the static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void detachTheUserProvidedFunctionAppFromTheStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .detachUserProvidedFunctionAppFromStaticSiteBuildWithResponse(
                "rg", "testStaticSite0", "12", "testFunctionApp", Context.NONE);
    }
}
```
