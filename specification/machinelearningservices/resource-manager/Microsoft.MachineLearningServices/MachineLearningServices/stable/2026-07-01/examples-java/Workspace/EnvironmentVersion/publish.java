
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for EnvironmentVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/EnvironmentVersion/publish.json
     */
    /**
     * Sample code: Publish Workspace Environment Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        publishWorkspaceEnvironmentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentVersions()
            .publish("test-rg", "my-aml-workspace", "string", "string", new DestinationAsset()
                .withDestinationName("string").withDestinationVersion("string").withRegistryName("string"),
                com.azure.core.util.Context.NONE);
    }
}
