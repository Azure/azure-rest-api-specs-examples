
import com.azure.resourcemanager.appservice.fluent.models.StaticSiteLinkedBackendArmResourceInner;

/**
 * Samples for StaticSites LinkBackendToBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/LinkBackendToStaticSiteBuild.json
     */
    /**
     * Sample code: Link a backend to a static site build.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void linkABackendToAStaticSiteBuild(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getStaticSites().linkBackendToBuild("rg", "testStaticSite0", "default", "testBackend",
            new StaticSiteLinkedBackendArmResourceInner().withBackendResourceId(
                "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/backendRg/providers/Microsoft.Web/sites/testBackend")
                .withRegion("West US 2"),
            com.azure.core.util.Context.NONE);
    }
}
