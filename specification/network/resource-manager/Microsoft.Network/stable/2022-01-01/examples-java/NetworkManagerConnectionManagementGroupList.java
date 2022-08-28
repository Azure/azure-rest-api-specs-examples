import com.azure.core.util.Context;

/** Samples for ManagementGroupNetworkManagerConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectionManagementGroupList.json
     */
    /**
     * Sample code: List Management Group Network Manager Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagementGroupNetworkManagerConnection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getManagementGroupNetworkManagerConnections()
            .list("managementGroupA", null, null, Context.NONE);
    }
}
