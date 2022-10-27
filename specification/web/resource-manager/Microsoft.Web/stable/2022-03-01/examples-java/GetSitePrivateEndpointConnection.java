import com.azure.core.util.Context;

/** Samples for WebApps GetPrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetSitePrivateEndpointConnection.json
     */
    /**
     * Sample code: Get a private endpoint connection for a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAPrivateEndpointConnectionForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getPrivateEndpointConnectionWithResponse("rg", "testSite", "connection", Context.NONE);
    }
}
