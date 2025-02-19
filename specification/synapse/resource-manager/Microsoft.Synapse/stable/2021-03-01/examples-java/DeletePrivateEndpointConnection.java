
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: Delete private endpoint connection.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void deletePrivateEndpointConnection(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateEndpointConnections().delete("ExampleResourceGroup", "ExampleWorkspace",
            "ExamplePrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
