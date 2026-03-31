
import com.azure.resourcemanager.appservice.models.StaticSitePatchResource;

/**
 * Samples for StaticSites UpdateStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/PatchStaticSite.json
     */
    /**
     * Sample code: Patch a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void patchAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().updateStaticSiteWithResponse("rg", "testStaticSite0",
            new StaticSitePatchResource(), com.azure.core.util.Context.NONE);
    }
}
