
/**
 * Samples for StaticSites ListStaticSiteUsers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteUsers.json
     */
    /**
     * Sample code: List users for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listUsersForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteUsers("rg", "testStaticSite0", "all",
            com.azure.core.util.Context.NONE);
    }
}
