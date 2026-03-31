
/**
 * Samples for StaticSites ListStaticSiteSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListStaticSiteSecrets.json
     */
    /**
     * Sample code: List secrets for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listSecretsForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().listStaticSiteSecretsWithResponse("rg", "testStaticSite0",
            com.azure.core.util.Context.NONE);
    }
}
