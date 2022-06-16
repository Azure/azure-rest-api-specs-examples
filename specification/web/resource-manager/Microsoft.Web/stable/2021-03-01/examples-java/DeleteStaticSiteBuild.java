import com.azure.core.util.Context;

/** Samples for StaticSites DeleteStaticSiteBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteStaticSiteBuild.json
     */
    /**
     * Sample code: Delete a static site build.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .deleteStaticSiteBuild("rg", "testStaticSite0", "12", Context.NONE);
    }
}
