/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get private endpoint connection.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getPrivateEndpointConnection(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "ExampleResourceGroup",
                "ExampleWorkspace",
                "ExamplePrivateEndpointConnection",
                com.azure.core.util.Context.NONE);
    }
}
