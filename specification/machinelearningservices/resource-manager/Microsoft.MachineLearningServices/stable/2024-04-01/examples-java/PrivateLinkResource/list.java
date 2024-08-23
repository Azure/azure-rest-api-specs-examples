
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/PrivateLinkResource/list.json
     */
    /**
     * Sample code: WorkspaceListPrivateLinkResources.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        workspaceListPrivateLinkResources(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.privateLinkResources().listWithResponse("rg-1234", "testworkspace", com.azure.core.util.Context.NONE);
    }
}
