
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteLinkedBackendArmResourceInner;

/**
 * Samples for StaticSites ValidateBackendForBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ValidateLinkedBackendForStaticSiteBuild.json
     */
    /**
     * Sample code: Validate if backend can be linked to static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        validateIfBackendCanBeLinkedToStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().validateBackendForBuild("rg", "testStaticSite0",
            "default", "testBackend",
            new StaticSiteLinkedBackendArmResourceInner().withBackendResourceId(
                "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/backendRg/providers/Microsoft.Web/sites/testBackend")
                .withRegion("West US 2"),
            com.azure.core.util.Context.NONE);
    }
}
