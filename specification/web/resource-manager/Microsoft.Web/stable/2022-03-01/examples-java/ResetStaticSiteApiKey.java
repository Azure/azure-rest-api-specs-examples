import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.StaticSiteResetPropertiesArmResource;

/** Samples for StaticSites ResetStaticSiteApiKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/ResetStaticSiteApiKey.json
     */
    /**
     * Sample code: Reset the api key for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetTheApiKeyForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .resetStaticSiteApiKeyWithResponse(
                "rg",
                "testStaticSite0",
                new StaticSiteResetPropertiesArmResource()
                    .withRepositoryToken("repoToken123")
                    .withShouldUpdateRepository(true),
                Context.NONE);
    }
}
