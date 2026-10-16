
/**
 * Samples for AdapterDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/GetAdapterDeployment.json
     */
    /**
     * Sample code: GetAdapterDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAdapterDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.adapterDeployments().getWithResponse("resourceGroupName", "accountName", "adapterDeploymentName",
            com.azure.core.util.Context.NONE);
    }
}
