
import com.azure.resourcemanager.machinelearning.models.ListViewType;

/**
 * Samples for FeaturesetVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetVersion/list.json
     */
    /**
     * Sample code: List Workspace Featureset Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceFeaturesetVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetVersions().list("test-rg", "my-aml-workspace", "string", null, "string", ListViewType.ALL,
            null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
