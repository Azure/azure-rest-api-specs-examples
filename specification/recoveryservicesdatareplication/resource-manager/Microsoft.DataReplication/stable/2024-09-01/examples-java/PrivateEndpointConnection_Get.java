
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnection_Get.json
     */
    /**
     * Sample code: Gets the Private Endpoint Connection.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsThePrivateEndpointConnection(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnections().getWithResponse("rgswagger_2024-09-01", "4", "vbkm",
            com.azure.core.util.Context.NONE);
    }
}
