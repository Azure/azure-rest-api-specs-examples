import com.azure.core.util.Context;

/** Samples for StaticSites ListStaticSiteUsers. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListStaticSiteUsers.json
     */
    /**
     * Sample code: List users for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listUsersForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteUsers("rg", "testStaticSite0", "all", Context.NONE);
    }
}
