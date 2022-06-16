import com.azure.core.util.Context;

/** Samples for StaticSites GetStaticSiteBuilds. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetStaticSiteBuilds.json
     */
    /**
     * Sample code: Get all builds for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllBuildsForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getStaticSiteBuilds("rg", "testStaticSite0", Context.NONE);
    }
}
