
/**
 * Samples for StaticSites GetLinkedBackend.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetLinkedBackendForStaticSite.json
     */
    /**
     * Sample code: Get details of the linked backend registered with a static site by name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsOfTheLinkedBackendRegisteredWithAStaticSiteByName(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getLinkedBackendWithResponse("rg", "testStaticSite0",
            "testBackend", com.azure.core.util.Context.NONE);
    }
}
