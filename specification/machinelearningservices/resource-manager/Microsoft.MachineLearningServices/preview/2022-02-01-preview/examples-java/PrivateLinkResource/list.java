import com.azure.core.util.Context;

/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/PrivateLinkResource/list.json
     */
    /**
     * Sample code: WorkspaceListPrivateLinkResources.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void workspaceListPrivateLinkResources(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.privateLinkResources().listWithResponse("rg-1234", "testworkspace", Context.NONE);
    }
}
