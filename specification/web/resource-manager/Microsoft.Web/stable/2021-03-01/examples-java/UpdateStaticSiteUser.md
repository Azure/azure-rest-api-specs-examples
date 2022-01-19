Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteUserArmResourceInner;

/** Samples for StaticSites UpdateStaticSiteUser. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/UpdateStaticSiteUser.json
     */
    /**
     * Sample code: Create or update a user for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAUserForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .updateStaticSiteUserWithResponse(
                "rg",
                "testStaticSite0",
                "aad",
                "1234",
                new StaticSiteUserArmResourceInner().withRoles("contributor"),
                Context.NONE);
    }
}
```
