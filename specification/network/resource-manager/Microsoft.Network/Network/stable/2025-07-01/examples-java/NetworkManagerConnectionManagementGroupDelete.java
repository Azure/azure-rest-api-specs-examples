
/**
 * Samples for ManagementGroupNetworkManagerConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectionManagementGroupDelete.json
     */
    /**
     * Sample code: Delete Management Group Network Manager Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteManagementGroupNetworkManagerConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getManagementGroupNetworkManagerConnections().deleteWithResponse("managementGroupA",
            "TestNMConnection", com.azure.core.util.Context.NONE);
    }
}
