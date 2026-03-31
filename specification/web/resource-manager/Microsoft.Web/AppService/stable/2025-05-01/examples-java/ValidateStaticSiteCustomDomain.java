
import com.azure.resourcemanager.appservice.models.StaticSiteCustomDomainRequestPropertiesArmResource;

/**
 * Samples for StaticSites ValidateCustomDomainCanBeAddedToStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ValidateStaticSiteCustomDomain.json
     */
    /**
     * Sample code: Validate a custom domain for a static site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        validateACustomDomainForAStaticSite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().validateCustomDomainCanBeAddedToStaticSite("rg", "testStaticSite0",
            "custom.domain.net", new StaticSiteCustomDomainRequestPropertiesArmResource(),
            com.azure.core.util.Context.NONE);
    }
}
