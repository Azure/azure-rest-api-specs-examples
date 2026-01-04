
/**
 * Samples for StaticSites GetUserProvidedFunctionAppsForStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetUserProvidedFunctionAppsForStaticSiteBuild.json
     */
    /**
     * Sample code: Get details of the user provided function apps registered with a static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppsRegisteredWithAStaticSiteBuild(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getUserProvidedFunctionAppsForStaticSiteBuild("rg",
            "testStaticSite0", "default", com.azure.core.util.Context.NONE);
    }
}
