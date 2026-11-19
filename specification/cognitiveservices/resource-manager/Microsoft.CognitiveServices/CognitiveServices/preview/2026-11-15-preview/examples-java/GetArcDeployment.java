
/**
 * Samples for ArcDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/GetArcDeployment.json
     */
    /**
     * Sample code: GetArcDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getArcDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.arcDeployments().getWithResponse("resourceGroupName", "accountName", "phi-3-gpu-arc",
            com.azure.core.util.Context.NONE);
    }
}
