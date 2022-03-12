Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites ListStaticSiteBuildFunctionAppSettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListStaticSiteBuildFunctionAppSettings.json
     */
    /**
     * Sample code: Get function app settings of a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFunctionAppSettingsOfAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteBuildFunctionAppSettingsWithResponse("rg", "testStaticSite0", "12", Context.NONE);
    }
}
```
