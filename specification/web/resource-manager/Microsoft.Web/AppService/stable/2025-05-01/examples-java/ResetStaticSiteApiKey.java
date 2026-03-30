
import com.azure.resourcemanager.appservice.models.StaticSiteResetPropertiesArmResource;

/**
 * Samples for StaticSites ResetStaticSiteApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ResetStaticSiteApiKey.json
     */
    /**
     * Sample code: Reset the api key for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void resetTheApiKeyForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().resetStaticSiteApiKeyWithResponse(
            "rg", "testStaticSite0", new StaticSiteResetPropertiesArmResource()
                .withRepositoryToken("fakeTokenPlaceholder").withShouldUpdateRepository(true),
            com.azure.core.util.Context.NONE);
    }
}
