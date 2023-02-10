/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateEndpointConnections.json
     */
    /**
     * Sample code: List private endpoint connections in workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listPrivateEndpointConnectionsInWorkspace(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateEndpointConnections()
            .list("ExampleResourceGroup", "ExampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
