
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
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
