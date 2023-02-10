/** Samples for Keys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteKey.json
     */
    /**
     * Sample code: Delete a workspace key.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAWorkspaceKey(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .keys()
            .deleteWithResponse(
                "ExampleResourceGroup", "ExampleWorkspace", "somekey", com.azure.core.util.Context.NONE);
    }
}
