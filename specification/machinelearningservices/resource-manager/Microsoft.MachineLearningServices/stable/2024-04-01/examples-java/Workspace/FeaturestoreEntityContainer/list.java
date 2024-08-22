
import com.azure.resourcemanager.machinelearning.models.ListViewType;

/**
 * Samples for FeaturestoreEntityContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturestoreEntityContainer/list.json
     */
    /**
     * Sample code: List Workspace Featurestore Entity Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listWorkspaceFeaturestoreEntityContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityContainers().list("test-rg", "my-aml-workspace", null, "string", ListViewType.ALL,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
