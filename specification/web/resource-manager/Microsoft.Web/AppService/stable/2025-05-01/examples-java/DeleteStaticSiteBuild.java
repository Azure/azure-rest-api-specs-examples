
/**
 * Samples for StaticSites DeleteStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteStaticSiteBuild.json
     */
    /**
     * Sample code: Delete a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().deleteStaticSiteBuild("rg", "testStaticSite0", "12",
            com.azure.core.util.Context.NONE);
    }
}
