Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites ListStaticSiteSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListStaticSiteSecrets.json
     */
    /**
     * Sample code: List secrets for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecretsForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteSecretsWithResponse("rg", "testStaticSite0", Context.NONE);
    }
}
```
