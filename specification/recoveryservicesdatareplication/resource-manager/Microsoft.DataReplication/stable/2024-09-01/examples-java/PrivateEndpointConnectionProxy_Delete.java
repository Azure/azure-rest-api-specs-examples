
/**
 * Samples for PrivateEndpointConnProxies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateEndpointConnectionProxy_Delete.json
     */
    /**
     * Sample code: Deletes the Private Endpoint Proxy Connection.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void deletesThePrivateEndpointProxyConnection(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateEndpointConnProxies().delete("rgswagger_2024-09-01", "4", "d", com.azure.core.util.Context.NONE);
    }
}
