
import com.azure.resourcemanager.appservice.models.StaticSiteCustomDomainRequestPropertiesArmResource;

/**
 * Samples for StaticSites ValidateCustomDomainCanBeAddedToStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ValidateStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Validate a custom domain for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validateACustomDomainForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().validateCustomDomainCanBeAddedToStaticSite("rg",
            "testStaticSite0", "custom.domain.net", new StaticSiteCustomDomainRequestPropertiesArmResource(),
            com.azure.core.util.Context.NONE);
    }
}
