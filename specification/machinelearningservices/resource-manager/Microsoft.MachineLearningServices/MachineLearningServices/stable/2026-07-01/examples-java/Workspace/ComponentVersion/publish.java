
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for ComponentVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentVersion/publish.json
     */
    /**
     * Sample code: Publish Workspace Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        publishWorkspaceComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentVersions()
            .publish("test-rg", "my-aml-workspace", "string", "string", new DestinationAsset()
                .withDestinationName("string").withDestinationVersion("string").withRegistryName("string"),
                com.azure.core.util.Context.NONE);
    }
}
