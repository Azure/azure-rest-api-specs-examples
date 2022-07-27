import com.azure.core.util.Context;

/** Samples for StaticSites GetPrivateEndpointConnectionList. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetSitePrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Get a list of private endpoint connections associated with a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAListOfPrivateEndpointConnectionsAssociatedWithASite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getPrivateEndpointConnectionList("rg", "testStaticSite0", Context.NONE);
    }
}
