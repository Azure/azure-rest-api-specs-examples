
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-06-16-preview/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: Gets the list of network security perimeter configurations of the private link scope.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getsTheListOfNetworkSecurityPerimeterConfigurationsOfThePrivateLinkScope(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.networkSecurityPerimeterConfigurations().listByPrivateLinkScope("my-resource-group",
            "my-privatelinkscope", com.azure.core.util.Context.NONE);
    }
}
