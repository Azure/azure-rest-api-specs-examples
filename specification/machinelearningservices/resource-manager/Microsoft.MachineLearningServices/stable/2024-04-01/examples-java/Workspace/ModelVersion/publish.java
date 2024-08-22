
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for ModelVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ModelVersion/publish.json
     */
    /**
     * Sample code: Publish Workspace Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        publishWorkspaceModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelVersions()
            .publish("test-rg", "my-aml-workspace", "string", "string", new DestinationAsset()
                .withRegistryName("string").withDestinationName("string").withDestinationVersion("string"),
                com.azure.core.util.Context.NONE);
    }
}
