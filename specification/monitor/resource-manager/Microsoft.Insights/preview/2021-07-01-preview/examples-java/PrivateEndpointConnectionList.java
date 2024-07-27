
/**
 * Samples for PrivateEndpointConnections ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a private link scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsListOfPrivateEndpointConnectionsOnAPrivateLinkScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateEndpointConnections()
            .listByPrivateLinkScopeWithResponse("MyResourceGroup", "MyPrivateLinkScope",
                com.azure.core.util.Context.NONE);
    }
}
