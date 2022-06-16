import com.azure.core.util.Context;

/** Samples for Keys Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetKey.json
     */
    /**
     * Sample code: Get a workspace key.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAWorkspaceKey(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.keys().getWithResponse("ExampleResourceGroup", "ExampleWorkspace", "somekey", Context.NONE);
    }
}
