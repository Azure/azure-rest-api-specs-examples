
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for CodeVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/CodeVersion/publish.json
     */
    /**
     * Sample code: Publish Workspace Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        publishWorkspaceCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions()
            .publish("test-rg", "my-aml-workspace", "string", "string", new DestinationAsset()
                .withDestinationName("string").withDestinationVersion("string").withRegistryName("string"),
                com.azure.core.util.Context.NONE);
    }
}
