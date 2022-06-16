import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.UnderlyingResourceAction;

/** Samples for Compute Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/delete.json
     */
    /**
     * Sample code: Delete Compute.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteCompute(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .computes()
            .delete("testrg123", "workspaces123", "compute123", UnderlyingResourceAction.DELETE, Context.NONE);
    }
}
