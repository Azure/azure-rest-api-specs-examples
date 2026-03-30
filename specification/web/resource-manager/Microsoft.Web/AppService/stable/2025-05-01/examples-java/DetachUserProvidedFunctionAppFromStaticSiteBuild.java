
/**
 * Samples for StaticSites DetachUserProvidedFunctionAppFromStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DetachUserProvidedFunctionAppFromStaticSiteBuild.json
     */
    /**
     * Sample code: Detach the user provided function app from the static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void detachTheUserProvidedFunctionAppFromTheStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().detachUserProvidedFunctionAppFromStaticSiteBuildWithResponse("rg",
            "testStaticSite0", "12", "testFunctionApp", com.azure.core.util.Context.NONE);
    }
}
