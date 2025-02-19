
import com.azure.resourcemanager.hardwaresecuritymodules.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.hardwaresecuritymodules.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.hardwaresecuritymodules.models.PrivateLinkServiceConnectionState;

/**
 * Samples for CloudHsmClusterPrivateEndpointConnections Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/CloudHsmClusterPrivateEndpointConnection_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudHsmClusterPrivateEndpointConnection_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void cloudHsmClusterPrivateEndpointConnectionCreateMaximumSetGen(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.cloudHsmClusterPrivateEndpointConnections().define("sample-pec")
            .withExistingCloudHsmCluster("rgcloudhsm", "chsm1")
            .withProperties(new PrivateEndpointConnectionProperties().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("My name is Joe and I'm approving this.")))
            .create();
    }
}
