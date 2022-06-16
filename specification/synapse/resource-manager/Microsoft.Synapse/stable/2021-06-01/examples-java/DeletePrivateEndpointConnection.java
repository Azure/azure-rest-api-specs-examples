import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: Delete private endpoint connection.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deletePrivateEndpointConnection(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateEndpointConnections()
            .delete("ExampleResourceGroup", "ExampleWorkspace", "ExamplePrivateEndpointConnection", Context.NONE);
    }
}
