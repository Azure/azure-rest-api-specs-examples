
import com.azure.resourcemanager.machinelearning.models.ListViewType;

/**
 * Samples for FeaturestoreEntityVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturestoreEntityVersion/list.json
     */
    /**
     * Sample code: List Workspace Featurestore Entity Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listWorkspaceFeaturestoreEntityVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityVersions().list("test-rg", "my-aml-workspace", "string", null, "string",
            ListViewType.ACTIVE_ONLY, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
