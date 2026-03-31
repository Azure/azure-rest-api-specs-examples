
/**
 * Samples for StaticSites UnlinkBackendFromBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UnlinkBackendFromStaticSiteBuild.json
     */
    /**
     * Sample code: Unlink a backend from a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        unlinkABackendFromAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().unlinkBackendFromBuildWithResponse("rg", "testStaticSite0", "12",
            "testBackend", null, com.azure.core.util.Context.NONE);
    }
}
