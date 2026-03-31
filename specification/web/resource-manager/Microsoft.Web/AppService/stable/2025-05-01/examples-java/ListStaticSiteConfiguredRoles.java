
/**
 * Samples for StaticSites ListStaticSiteConfiguredRoles.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteConfiguredRoles.json
     */
    /**
     * Sample code: Lists the configured roles for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listsTheConfiguredRolesForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteConfiguredRolesWithResponse("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
