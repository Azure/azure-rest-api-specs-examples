
/**
 * Samples for StaticSites GetStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSiteBuild.json
     */
    /**
     * Sample code: Get a static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getStaticSiteBuildWithResponse("rg",
            "testStaticSite0", "12", com.azure.core.util.Context.NONE);
    }
}
