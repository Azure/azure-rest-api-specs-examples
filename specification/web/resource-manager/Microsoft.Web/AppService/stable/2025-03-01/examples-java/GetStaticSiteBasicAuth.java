
import com.azure.resourcemanager.appservice.models.BasicAuthName;

/**
 * Samples for StaticSites GetBasicAuth.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSiteBasicAuth.
     * json
     */
    /**
     * Sample code: Gets the basic auth properties for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheBasicAuthPropertiesForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().getBasicAuthWithResponse("rg", "testStaticSite0",
            BasicAuthName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
