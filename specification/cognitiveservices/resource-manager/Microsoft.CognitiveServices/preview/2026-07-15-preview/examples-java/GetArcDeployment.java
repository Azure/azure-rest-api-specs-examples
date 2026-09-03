
/**
 * Samples for ArcDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15-preview/GetArcDeployment.json
     */
    /**
     * Sample code: GetArcDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getArcDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.arcDeployments().getWithResponse("resourceGroupName", "accountName", "qwen-template-arc",
            com.azure.core.util.Context.NONE);
    }
}
