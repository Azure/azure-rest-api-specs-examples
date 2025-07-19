
/**
 * Samples for PrivateEndpointConnections ListByCloudHsmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/CloudHsmClusterPrivateEndpointConnection_ListByCloudHsmCluster_MaximumSet_Gen.json
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
