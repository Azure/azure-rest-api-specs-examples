
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteBasicAuthPropertiesArmResourceInner;
import com.azure.resourcemanager.appservice.models.BasicAuthName;

/**
 * Samples for StaticSites CreateOrUpdateBasicAuth.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * CreateOrUpdateStaticSiteBasicAuth.json
     */
    /**
     * Sample code: Creates or updates basic auth properties for a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createsOrUpdatesBasicAuthPropertiesForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().createOrUpdateBasicAuthWithResponse("rg",
            "testStaticSite0", BasicAuthName.DEFAULT, new StaticSiteBasicAuthPropertiesArmResourceInner()
                .withPassword("fakeTokenPlaceholder").withApplicableEnvironmentsMode("AllEnvironments"),
            com.azure.core.util.Context.NONE);
    }
}
