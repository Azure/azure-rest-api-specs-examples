
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PrivateEndpointConnectionsList.json
     */
    /**
     * Sample code: List Private Endpoint Connections by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listPrivateEndpointConnectionsByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().listWithResponse("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
