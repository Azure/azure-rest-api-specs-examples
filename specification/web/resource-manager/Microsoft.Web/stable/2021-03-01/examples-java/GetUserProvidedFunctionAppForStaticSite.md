Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites GetUserProvidedFunctionAppForStaticSite. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetUserProvidedFunctionAppForStaticSite.json
     */
    /**
     * Sample code: Get details of the user provided function app registered with a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getUserProvidedFunctionAppForStaticSiteWithResponse(
                "rg", "testStaticSite0", "testFunctionApp", Context.NONE);
    }
}
```
