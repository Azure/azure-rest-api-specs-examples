
/**
 * Samples for Projects Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetProject.json
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
