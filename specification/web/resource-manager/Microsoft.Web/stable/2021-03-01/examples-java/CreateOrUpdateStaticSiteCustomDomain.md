Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.StaticSiteCustomDomainRequestPropertiesArmResource;

/** Samples for StaticSites CreateOrUpdateStaticSiteCustomDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Create or update a custom domain for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateACustomDomainForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .createOrUpdateStaticSiteCustomDomain(
                "rg",
                "testStaticSite0",
                "custom.domain.net",
                new StaticSiteCustomDomainRequestPropertiesArmResource(),
                Context.NONE);
    }
}
```
