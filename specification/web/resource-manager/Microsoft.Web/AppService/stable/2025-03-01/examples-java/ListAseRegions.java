
/**
 * Samples for ResourceProvider ListAseRegions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListAseRegions.json
     */
    /**
     * Sample code: List aseregions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAseregions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getResourceProviders()
            .listAseRegions(com.azure.core.util.Context.NONE);
    }
}
