
/**
 * Samples for CloudHsmClusterPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/CloudHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterPrivateEndpointConnectionGetMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusterPrivateEndpointConnections().getWithResponse("rgcloudhsm", "chsm1", "sample-pec",
            com.azure.core.util.Context.NONE);
    }
}
