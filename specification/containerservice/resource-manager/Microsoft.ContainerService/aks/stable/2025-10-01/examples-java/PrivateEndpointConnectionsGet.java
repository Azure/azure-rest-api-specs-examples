
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: Get Private Endpoint Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getPrivateEndpointConnections().getWithResponse("rg1",
            "clustername1", "privateendpointconnection1", com.azure.core.util.Context.NONE);
    }
}
