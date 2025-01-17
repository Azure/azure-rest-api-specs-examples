
/**
 * Samples for IntegrationRuntimeObjectMetadata Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimeObjectMetadata_Refresh.json
     */
    /**
     * Sample code: Refresh object metadata.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void refreshObjectMetadata(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeObjectMetadatas().refresh("exampleResourceGroup", "exampleWorkspace",
            "testactivityv2", com.azure.core.util.Context.NONE);
    }
}
