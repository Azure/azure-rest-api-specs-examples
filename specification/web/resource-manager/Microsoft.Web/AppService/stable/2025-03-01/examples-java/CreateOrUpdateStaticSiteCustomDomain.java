
import com.azure.resourcemanager.appservice.models.StaticSiteCustomDomainRequestPropertiesArmResource;

/**
 * Samples for StaticSites CreateOrUpdateStaticSiteCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * CreateOrUpdateStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Create or update a custom domain for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateACustomDomainForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().createOrUpdateStaticSiteCustomDomain("rg",
            "testStaticSite0", "custom.domain.net", new StaticSiteCustomDomainRequestPropertiesArmResource(),
            com.azure.core.util.Context.NONE);
    }
}
