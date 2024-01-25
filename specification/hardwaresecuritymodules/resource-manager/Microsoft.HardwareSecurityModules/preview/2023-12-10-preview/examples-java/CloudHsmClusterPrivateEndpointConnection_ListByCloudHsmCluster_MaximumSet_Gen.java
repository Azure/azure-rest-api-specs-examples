
/**
 * Samples for PrivateEndpointConnections ListByCloudHsmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2023-12-10-
     * preview/examples/CloudHsmClusterPrivateEndpointConnection_ListByCloudHsmCluster_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmClusterPrivateEndpointConnection_ListByResource_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterPrivateEndpointConnectionListByResourceMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.privateEndpointConnections().listByCloudHsmCluster("rgcloudhsm", "chsm1",
            com.azure.core.util.Context.NONE);
    }
}
