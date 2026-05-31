
/**
 * Samples for ManagedComputeDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteManagedComputeDeployment.json
     */
    /**
     * Sample code: DeleteManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteManagedComputeDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeDeployments().delete("resourceGroupName", "accountName", "gpt-oss-120b-gpu",
            com.azure.core.util.Context.NONE);
    }
}
