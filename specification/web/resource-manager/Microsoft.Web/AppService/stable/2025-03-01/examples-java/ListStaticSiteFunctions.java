
/**
 * Samples for StaticSites ListStaticSiteFunctions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListStaticSiteFunctions.
     * json
     */
    /**
     * Sample code: Gets the functions of a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheFunctionsOfAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().listStaticSiteFunctions("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
