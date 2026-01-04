
/**
 * Samples for StaticSites GetUserProvidedFunctionAppForStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetUserProvidedFunctionAppForStaticSite.json
     */
    /**
     * Sample code: Get details of the user provided function app registered with a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getUserProvidedFunctionAppForStaticSiteWithResponse(
            "rg", "testStaticSite0", "testFunctionApp", com.azure.core.util.Context.NONE);
    }
}
