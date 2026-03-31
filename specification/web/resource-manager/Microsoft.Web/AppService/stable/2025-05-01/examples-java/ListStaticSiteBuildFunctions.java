
/**
 * Samples for StaticSites ListStaticSiteBuildFunctions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteBuildFunctions.json
     */
    /**
     * Sample code: Gets the functions of a particular static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getsTheFunctionsOfAParticularStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteBuildFunctions("rg", "testStaticSite0", "default",
            com.azure.core.util.Context.NONE);
    }
}
