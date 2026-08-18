
/**
 * Samples for NetworkSecurityPerimeterConfigurations ReconcileForPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-07-15/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationReconcile.json
     */
    /**
     * Sample code: Reconciles the network security perimeter configuration of the private link scope.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void reconcilesTheNetworkSecurityPerimeterConfigurationOfThePrivateLinkScope(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.networkSecurityPerimeterConfigurations().reconcileForPrivateLinkScope("my-resource-group",
            "my-privatelinkscope", "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation",
            com.azure.core.util.Context.NONE);
    }
}
