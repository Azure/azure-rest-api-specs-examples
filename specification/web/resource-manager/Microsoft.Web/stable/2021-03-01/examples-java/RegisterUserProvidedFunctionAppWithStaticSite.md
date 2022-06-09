```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteUserProvidedFunctionAppArmResourceInner;

/** Samples for StaticSites RegisterUserProvidedFunctionAppWithStaticSite. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/RegisterUserProvidedFunctionAppWithStaticSite.json
     */
    /**
     * Sample code: Register a user provided function app with a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registerAUserProvidedFunctionAppWithAStaticSite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .registerUserProvidedFunctionAppWithStaticSite(
                "rg",
                "testStaticSite0",
                "testFunctionApp",
                new StaticSiteUserProvidedFunctionAppArmResourceInner()
                    .withFunctionAppResourceId(
                        "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp")
                    .withFunctionAppRegion("West US 2"),
                true,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
