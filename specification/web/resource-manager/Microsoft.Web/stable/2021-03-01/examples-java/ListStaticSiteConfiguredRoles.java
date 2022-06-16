import com.azure.core.util.Context;

/** Samples for StaticSites ListStaticSiteConfiguredRoles. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListStaticSiteConfiguredRoles.json
     */
    /**
     * Sample code: Lists the configured roles for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsTheConfiguredRolesForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteConfiguredRolesWithResponse("rg", "testStaticSite0", Context.NONE);
    }
}
