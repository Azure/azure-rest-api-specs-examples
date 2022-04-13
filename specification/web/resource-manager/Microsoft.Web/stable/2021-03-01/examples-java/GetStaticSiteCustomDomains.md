Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites ListStaticSiteCustomDomains. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetStaticSiteCustomDomains.json
     */
    /**
     * Sample code: List custom domains for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCustomDomainsForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteCustomDomains("rg", "testStaticSite0", Context.NONE);
    }
}
```
