
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateEndpointConnections().getWithResponse(
            "MyResourceGroup", "MyPrivateLinkScope", "private-endpoint-connection-name",
            com.azure.core.util.Context.NONE);
    }
}
