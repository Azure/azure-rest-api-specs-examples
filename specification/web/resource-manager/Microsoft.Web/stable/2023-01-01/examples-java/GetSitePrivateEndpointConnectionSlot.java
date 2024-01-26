
/**
 * Samples for WebApps GetPrivateEndpointConnectionSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetSitePrivateEndpointConnectionSlot.
     * json
     */
    /**
     * Sample code: Get a private endpoint connection for a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAPrivateEndpointConnectionForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getPrivateEndpointConnectionSlotWithResponse("rg",
            "testSite", "connection", "stage", com.azure.core.util.Context.NONE);
    }
}
