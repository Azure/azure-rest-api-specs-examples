import com.azure.core.util.Context;

/** Samples for StaticSites DetachUserProvidedFunctionAppFromStaticSite. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/DetachUserProvidedFunctionAppFromStaticSite.json
     */
    /**
     * Sample code: Detach the user provided function app from the static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void detachTheUserProvidedFunctionAppFromTheStaticSite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .detachUserProvidedFunctionAppFromStaticSiteWithResponse(
                "rg", "testStaticSite0", "testFunctionApp", Context.NONE);
    }
}
