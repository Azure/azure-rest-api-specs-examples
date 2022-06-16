import com.azure.core.util.Context;

/** Samples for WebApps DeletePrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteSitePrivateEndpointConnection.json
     */
    /**
     * Sample code: Delete a private endpoint connection for a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPrivateEndpointConnectionForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .deletePrivateEndpointConnection("rg", "testSite", "connection", Context.NONE);
    }
}
