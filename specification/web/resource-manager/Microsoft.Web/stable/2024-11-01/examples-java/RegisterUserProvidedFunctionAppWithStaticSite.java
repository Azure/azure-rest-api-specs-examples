
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteUserProvidedFunctionAppArmResourceInner;

/**
 * Samples for StaticSites RegisterUserProvidedFunctionAppWithStaticSite.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/
     * RegisterUserProvidedFunctionAppWithStaticSite.json
     */
    /**
     * Sample code: Register a user provided function app with a static site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        registerAUserProvidedFunctionAppWithAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().registerUserProvidedFunctionAppWithStaticSite("rg",
            "testStaticSite0", "testFunctionApp",
            new StaticSiteUserProvidedFunctionAppArmResourceInner().withFunctionAppResourceId(
                "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp")
                .withFunctionAppRegion("West US 2"),
            true, com.azure.core.util.Context.NONE);
    }
}
