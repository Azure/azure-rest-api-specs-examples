
/**
 * Samples for StaticSites UnlinkBackendFromBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * UnlinkBackendFromStaticSiteBuild.json
     */
    /**
     * Sample code: Unlink a backend from a static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void unlinkABackendFromAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().unlinkBackendFromBuildWithResponse("rg",
            "testStaticSite0", "12", "testBackend", null, com.azure.core.util.Context.NONE);
    }
}
