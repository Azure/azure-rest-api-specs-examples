
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/
     * NetworkSecurityPerimeterConfigurationsListByService.json
     */
    /**
     * Sample code: List NSP configs by search service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNSPConfigsBySearchService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getNetworkSecurityPerimeterConfigurations()
            .listByService("rg1", "mysearchservice", com.azure.core.util.Context.NONE);
    }
}
