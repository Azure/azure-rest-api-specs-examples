
import com.azure.resourcemanager.appservice.models.BasicAuthName;

/**
 * Samples for StaticSites GetBasicAuth.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetStaticSiteBasicAuth.json
     */
    /**
     * Sample code: Gets the basic auth properties for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getsTheBasicAuthPropertiesForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().getBasicAuthWithResponse("rg", "testStaticSite0",
            BasicAuthName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
