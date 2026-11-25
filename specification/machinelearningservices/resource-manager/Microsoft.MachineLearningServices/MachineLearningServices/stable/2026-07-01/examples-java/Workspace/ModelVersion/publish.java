
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for ModelVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ModelVersion/publish.json
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
                .withDestinationName("string").withDestinationVersion("string").withRegistryName("string"),
                com.azure.core.util.Context.NONE);
    }
}
