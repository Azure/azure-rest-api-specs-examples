
/**
 * Samples for StaticSites ListStaticSiteSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListStaticSiteSecrets.json
     */
    /**
     * Sample code: List secrets for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecretsForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().listStaticSiteSecretsWithResponse("rg",
            "testStaticSite0", com.azure.core.util.Context.NONE);
    }
}
