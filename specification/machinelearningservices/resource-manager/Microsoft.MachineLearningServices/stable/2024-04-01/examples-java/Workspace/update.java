
import com.azure.resourcemanager.machinelearning.models.PublicNetworkAccess;
import com.azure.resourcemanager.machinelearning.models.Workspace;

/**
 * Samples for Workspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/update.json
     */
    /**
     * Sample code: Update Workspace.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void updateWorkspace(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        Workspace resource = manager.workspaces()
            .getByResourceGroupWithResponse("workspace-1234", "testworkspace", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withDescription("new description").withFriendlyName("New friendly name")
            .withPublicNetworkAccess(PublicNetworkAccess.DISABLED).apply();
    }
}
