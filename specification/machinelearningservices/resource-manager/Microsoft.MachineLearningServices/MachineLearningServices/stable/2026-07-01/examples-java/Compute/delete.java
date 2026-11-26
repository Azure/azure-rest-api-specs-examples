
import com.azure.resourcemanager.machinelearning.models.UnderlyingResourceAction;

/**
 * Samples for Compute Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Compute/delete.json
     */
    /**
     * Sample code: Delete Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().delete("testrg123", "workspaces123", "compute123", UnderlyingResourceAction.DELETE,
            com.azure.core.util.Context.NONE);
    }
}
