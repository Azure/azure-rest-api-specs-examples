
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * GetProject.json
     */
    /**
     * Sample code: Get Project.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getProject(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().getWithResponse("myResourceGroup", "myAccount", "myProject",
            com.azure.core.util.Context.NONE);
    }
}
