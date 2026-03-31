
/**
 * Samples for StaticSites GetLinkedBackendForBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetLinkedBackendForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the linked backend registered with a static site build by name.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getDetailsOfTheLinkedBackendRegisteredWithAStaticSiteBuildByName(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getLinkedBackendForBuildWithResponse("rg", "testStaticSite0",
            "default", "testBackend", com.azure.core.util.Context.NONE);
    }
}
