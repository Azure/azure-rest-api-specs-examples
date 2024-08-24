
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/GetPrivateEndpointConnection.
     * json
     */
    /**
     * Sample code: PrivateEndpointConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndpointConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getPrivateEndpointConnections().getWithResponse("rg1",
            "mysearchservice", "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546", null,
            com.azure.core.util.Context.NONE);
    }
}
