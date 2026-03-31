
/**
 * Samples for StaticSites GetStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBuild.json
     */
    /**
     * Sample code: Get a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getStaticSiteBuildWithResponse("rg", "testStaticSite0", "12",
            com.azure.core.util.Context.NONE);
    }
}
