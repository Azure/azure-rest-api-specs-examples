
/**
 * Samples for StaticSites GetLinkedBackendForBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetLinkedBackendForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the linked backend registered with a static site build by name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheLinkedBackendRegisteredWithAStaticSiteBuildByName(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getLinkedBackendForBuildWithResponse("rg",
            "testStaticSite0", "default", "testBackend", com.azure.core.util.Context.NONE);
    }
}
