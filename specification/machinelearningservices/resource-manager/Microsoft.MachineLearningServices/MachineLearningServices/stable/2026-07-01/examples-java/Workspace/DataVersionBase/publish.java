
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for DataVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/DataVersionBase/publish.json
     */
    /**
     * Sample code: Publish Workspace Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        publishWorkspaceDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataVersions()
            .publish("test-rg", "my-aml-workspace", "string", "string", new DestinationAsset()
                .withDestinationName("string").withDestinationVersion("string").withRegistryName("string"),
                com.azure.core.util.Context.NONE);
    }
}
