Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteArmResourceInner;
import com.azure.resourcemanager.appservice.models.SkuDescription;
import com.azure.resourcemanager.appservice.models.StaticSiteBuildProperties;

/** Samples for StaticSites CreateOrUpdateStaticSite. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateStaticSite.json
     */
    /**
     * Sample code: Create or update a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .createOrUpdateStaticSite(
                "rg",
                "testStaticSite0",
                new StaticSiteArmResourceInner()
                    .withLocation("West US 2")
                    .withSku(new SkuDescription().withName("Basic").withTier("Basic"))
                    .withRepositoryUrl("https://github.com/username/RepoName")
                    .withBranch("master")
                    .withRepositoryToken("repoToken123")
                    .withBuildProperties(
                        new StaticSiteBuildProperties()
                            .withAppLocation("app")
                            .withApiLocation("api")
                            .withAppArtifactLocation("build")),
                Context.NONE);
    }
}
```
