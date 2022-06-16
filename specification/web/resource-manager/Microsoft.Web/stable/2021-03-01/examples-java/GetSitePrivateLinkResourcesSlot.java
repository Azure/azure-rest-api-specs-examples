import com.azure.core.util.Context;

/** Samples for WebApps GetPrivateLinkResourcesSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSitePrivateLinkResourcesSlot.json
     */
    /**
     * Sample code: Get private link resources of a site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateLinkResourcesOfASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getPrivateLinkResourcesSlotWithResponse("rg", "testSite", "stage", Context.NONE);
    }
}
