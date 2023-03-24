import com.azure.resourcemanager.appservice.fluent.models.StaticSiteLinkedBackendArmResourceInner;

/** Samples for StaticSites LinkBackendToBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/LinkBackendToStaticSiteBuild.json
     */
    /**
     * Sample code: Link a backend to a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void linkABackendToAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .linkBackendToBuild(
                "rg",
                "testStaticSite0",
                "default",
                "testBackend",
                new StaticSiteLinkedBackendArmResourceInner()
                    .withBackendResourceId(
                        "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/backendRg/providers/Microsoft.Web/sites/testBackend")
                    .withRegion("West US 2"),
                com.azure.core.util.Context.NONE);
    }
}
