import com.azure.core.util.Context;

/** Samples for StaticSites Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteStaticSite.json
     */
    /**
     * Sample code: Delete a static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAStaticSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().delete("rg", "testStaticSite0", Context.NONE);
    }
}
