
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/NetworkSecurityPerimeterConfigurationsListByService.json
     */
    /**
     * Sample code: List NSP configs by search service.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void listNSPConfigsBySearchService(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().listByService("rg1", "mysearchservice",
            com.azure.core.util.Context.NONE);
    }
}
