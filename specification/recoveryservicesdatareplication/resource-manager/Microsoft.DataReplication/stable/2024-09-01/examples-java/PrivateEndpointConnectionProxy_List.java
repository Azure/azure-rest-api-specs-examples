
/**
 * Samples for PrivateEndpointConnProxies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnectionProxy_List.json
     */
    /**
     * Sample code: Lists the Private Endpoint Connection Proxy.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsThePrivateEndpointConnectionProxy(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnProxies().list("rgswagger_2024-09-01", "4", com.azure.core.util.Context.NONE);
    }
}
