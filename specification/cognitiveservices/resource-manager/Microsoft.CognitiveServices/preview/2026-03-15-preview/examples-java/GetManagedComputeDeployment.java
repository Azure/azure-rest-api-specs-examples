
/**
 * Samples for ManagedComputeDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetManagedComputeDeployment.json
     */
    /**
     * Sample code: GetManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getManagedComputeDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeDeployments().getWithResponse("resourceGroupName", "accountName", "gpt-oss-120b-gpu",
            com.azure.core.util.Context.NONE);
    }
}
