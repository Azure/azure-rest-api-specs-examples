
/**
 * Samples for NetworkSecurityPerimeterConfigurations ReconcileForPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/
     * networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationReconcile.json
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
