
/**
 * Samples for StaticSites ListStaticSiteFunctionAppSettings.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListStaticSiteFunctionAppSettings.json
     */
    /**
     * Sample code: Get function app settings of a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFunctionAppSettingsOfAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().listStaticSiteFunctionAppSettingsWithResponse("rg",
            "testStaticSite0", com.azure.core.util.Context.NONE);
    }
}
