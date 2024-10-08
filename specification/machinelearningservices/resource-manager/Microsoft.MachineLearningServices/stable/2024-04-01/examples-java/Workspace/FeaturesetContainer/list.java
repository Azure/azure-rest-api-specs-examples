
import com.azure.resourcemanager.machinelearning.models.ListViewType;

/**
 * Samples for FeaturesetContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetContainer/list.json
     */
    /**
     * Sample code: List Workspace Featureset Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceFeaturesetContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetContainers().list("test-rg", "my-aml-workspace", null, "string", ListViewType.ARCHIVED_ONLY,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
