
/**
 * Samples for ManagementGroupNetworkManagerConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerConnectionManagementGroupGet.json
     */
    /**
     * Sample code: Get Management Group Network Manager Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getManagementGroupNetworkManagerConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getManagementGroupNetworkManagerConnections()
            .getWithResponse("managementGroupA", "TestNMConnection", com.azure.core.util.Context.NONE);
    }
}
