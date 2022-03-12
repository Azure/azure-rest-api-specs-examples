Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites GetUserProvidedFunctionAppForStaticSiteBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetUserProvidedFunctionAppForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the user provided function app registered with a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getUserProvidedFunctionAppForStaticSiteBuildWithResponse(
                "rg", "testStaticSite0", "default", "testFunctionApp", Context.NONE);
    }
}
```
