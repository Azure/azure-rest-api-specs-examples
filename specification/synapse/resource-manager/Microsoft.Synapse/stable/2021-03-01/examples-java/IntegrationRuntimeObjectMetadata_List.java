
import com.azure.resourcemanager.synapse.models.GetSsisObjectMetadataRequest;

/**
 * Samples for IntegrationRuntimeObjectMetadata List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimeObjectMetadata_List.json
     */
    /**
     * Sample code: Get integration runtime object metadata.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getIntegrationRuntimeObjectMetadata(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeObjectMetadatas().listWithResponse("exampleResourceGroup", "exampleWorkspace",
            "testactivityv2", new GetSsisObjectMetadataRequest().withMetadataPath("ssisFolders"),
            com.azure.core.util.Context.NONE);
    }
}
