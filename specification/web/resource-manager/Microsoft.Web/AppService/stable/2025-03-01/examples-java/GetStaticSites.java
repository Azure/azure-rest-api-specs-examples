
/**
 * Samples for StaticSites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSites.json
     */
    /**
     * Sample code: Get static sites for a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getStaticSitesForAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().listByResourceGroup("rg",
            com.azure.core.util.Context.NONE);
    }
}
