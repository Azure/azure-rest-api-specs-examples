
/**
 * Samples for StaticSites ListStaticSiteAppSettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListStaticSiteAppSettings.
     * json
     */
    /**
     * Sample code: Get app settings of a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppSettingsOfAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().listStaticSiteAppSettingsWithResponse("rg",
            "testStaticSite0", com.azure.core.util.Context.NONE);
    }
}
