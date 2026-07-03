
/**
 * Samples for ManagementGroupNetworkManagerConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerConnectionManagementGroupList.json
     */
    /**
     * Sample code: List Management Group Network Manager Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listManagementGroupNetworkManagerConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getManagementGroupNetworkManagerConnections().list("managementGroupA", null, null,
            com.azure.core.util.Context.NONE);
    }
}
