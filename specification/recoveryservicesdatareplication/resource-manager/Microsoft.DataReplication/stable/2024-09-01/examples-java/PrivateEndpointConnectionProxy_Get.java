
/**
 * Samples for PrivateEndpointConnProxies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnectionProxy_Get.json
     */
    /**
     * Sample code: Get Private Endpoint Connection Proxy.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getPrivateEndpointConnectionProxy(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnProxies().getWithResponse("rgswagger_2024-09-01", "4", "d",
            com.azure.core.util.Context.NONE);
    }
}
