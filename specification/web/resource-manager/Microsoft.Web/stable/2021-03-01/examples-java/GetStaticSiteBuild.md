Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StaticSites GetStaticSiteBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetStaticSiteBuild.json
     */
    /**
     * Sample code: Get a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getStaticSiteBuildWithResponse("rg", "testStaticSite0", "12", Context.NONE);
    }
}
```
