
/**
 * Samples for StaticSites DetachUserProvidedFunctionAppFromStaticSiteBuild.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * DetachUserProvidedFunctionAppFromStaticSiteBuild.json
     */
    /**
     * Sample code: Detach the user provided function app from the static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        detachTheUserProvidedFunctionAppFromTheStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites()
            .detachUserProvidedFunctionAppFromStaticSiteBuildWithResponse("rg", "testStaticSite0", "12",
                "testFunctionApp", com.azure.core.util.Context.NONE);
    }
}
