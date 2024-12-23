
/**
 * Samples for BigDataPools ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/ListBigDataPoolsInWorkspace.
     * json
     */
    /**
     * Sample code: List Big Data pools in a workspace.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void listBigDataPoolsInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.bigDataPools().listByWorkspace("ExampleResourceGroup", "ExampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
