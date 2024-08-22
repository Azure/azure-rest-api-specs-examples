
import com.azure.resourcemanager.machinelearning.models.DestinationAsset;

/**
 * Samples for DataVersions Publish.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/DataVersionBase/publish.json
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
                .withRegistryName("string").withDestinationName("string").withDestinationVersion("string"),
                com.azure.core.util.Context.NONE);
    }
}
