
/**
 * Samples for ManagedComputeDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetVmManagedComputeDeployment.json
     */
    /**
     * Sample code: GetVmManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getVmManagedComputeDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeDeployments().getWithResponse("resourceGroupName", "accountName", "gpt-oss-120b-byoc",
            com.azure.core.util.Context.NONE);
    }
}
