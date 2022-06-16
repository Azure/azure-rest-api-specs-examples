import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.PublicNetworkAccess;
import com.azure.resourcemanager.machinelearning.models.Workspace;

/** Samples for Workspaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/update.json
     */
    /**
     * Sample code: Update Workspace.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void updateWorkspace(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        Workspace resource =
            manager
                .workspaces()
                .getByResourceGroupWithResponse("workspace-1234", "testworkspace", Context.NONE)
                .getValue();
        resource
            .update()
            .withDescription("new description")
            .withFriendlyName("New friendly name")
            .withPublicNetworkAccess(PublicNetworkAccess.DISABLED)
            .apply();
    }
}
