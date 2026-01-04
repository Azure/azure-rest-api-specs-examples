
/**
 * Samples for StaticSites ListBasicAuth.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListStaticSiteBasicAuth.
     * json
     */
    /**
     * Sample code: Lists the basic auth properties for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsTheBasicAuthPropertiesForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().listBasicAuth("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
