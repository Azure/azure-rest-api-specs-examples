/** Samples for StaticSites ListStaticSiteBuildFunctionAppSettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/
     * ListStaticSiteBuildFunctionAppSettings.json
     */
    /**
     * Sample code: Get function app settings of a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFunctionAppSettingsOfAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .listStaticSiteBuildFunctionAppSettingsWithResponse(
                "rg", "testStaticSite0", "12", com.azure.core.util.Context.NONE);
    }
}
