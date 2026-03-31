
/**
 * Samples for StaticSites DeleteStaticSiteUser.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteStaticSiteUser.json
     */
    /**
     * Sample code: Delete a user for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteAUserForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().deleteStaticSiteUserWithResponse("rg", "testStaticSite0", "aad",
            "1234", com.azure.core.util.Context.NONE);
    }
}
