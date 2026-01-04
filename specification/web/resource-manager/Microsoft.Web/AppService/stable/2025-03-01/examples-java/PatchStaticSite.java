
import com.azure.resourcemanager.appservice.models.StaticSitePatchResource;

/**
 * Samples for StaticSites UpdateStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/PatchStaticSite.json
     */
    /**
     * Sample code: Patch a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().updateStaticSiteWithResponse("rg", "testStaticSite0",
            new StaticSitePatchResource(), com.azure.core.util.Context.NONE);
    }
}
