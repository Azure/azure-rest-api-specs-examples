import com.azure.core.util.Context;

/** Samples for StaticSites GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetStaticSite.json
     */
    /**
     * Sample code: Get details for a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDetailsForAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .getByResourceGroupWithResponse("rg", "testStaticSite0", Context.NONE);
    }
}
