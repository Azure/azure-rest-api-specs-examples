
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteUserProvidedFunctionAppArmResourceInner;

/**
 * Samples for StaticSites RegisterUserProvidedFunctionAppWithStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/RegisterUserProvidedFunctionAppWithStaticSiteBuild.json
     */
    /**
     * Sample code: Register a user provided function app with a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void registerAUserProvidedFunctionAppWithAStaticSiteBuild(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().registerUserProvidedFunctionAppWithStaticSiteBuild("rg",
            "testStaticSite0", "default", "testFunctionApp",
            new StaticSiteUserProvidedFunctionAppArmResourceInner().withFunctionAppResourceId(
                "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp")
                .withFunctionAppRegion("West US 2"),
            true, com.azure.core.util.Context.NONE);
    }
}
