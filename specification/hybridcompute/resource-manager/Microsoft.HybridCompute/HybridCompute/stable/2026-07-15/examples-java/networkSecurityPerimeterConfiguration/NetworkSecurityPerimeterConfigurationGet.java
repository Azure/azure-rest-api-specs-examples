
/**
 * Samples for NetworkSecurityPerimeterConfigurations GetByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-07-15/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationGet.json
     */
    /**
     * Sample code: Gets the network security perimeter configuration of the private link scope.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getsTheNetworkSecurityPerimeterConfigurationOfThePrivateLinkScope(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.networkSecurityPerimeterConfigurations().getByPrivateLinkScopeWithResponse("my-resource-group",
            "my-privatelinkscope", "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation",
            com.azure.core.util.Context.NONE);
    }
}
