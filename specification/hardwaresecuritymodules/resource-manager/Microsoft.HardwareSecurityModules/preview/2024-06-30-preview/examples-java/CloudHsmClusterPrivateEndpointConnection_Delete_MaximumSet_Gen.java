
/**
 * Samples for CloudHsmClusterPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmClusterPrivateEndpointConnection_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmClusterPrivateEndpointConnection_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterPrivateEndpointConnectionDeleteMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusterPrivateEndpointConnections().delete("rgcloudhsm", "chsm1", "sample-pec",
            com.azure.core.util.Context.NONE);
    }
}
