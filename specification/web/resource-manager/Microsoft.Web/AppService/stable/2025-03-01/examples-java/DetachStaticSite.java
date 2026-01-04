
/**
 * Samples for StaticSites DetachStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/DetachStaticSite.json
     */
    /**
     * Sample code: Detach a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void detachAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().detachStaticSite("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
