/** Samples for Keys ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListKeysInWorkspace.json
     */
    /**
     * Sample code: List keys in workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listKeysInWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.keys().listByWorkspace("ExampleResourceGroup", "ExampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
