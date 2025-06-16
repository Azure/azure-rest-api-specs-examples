
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationList.json
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
