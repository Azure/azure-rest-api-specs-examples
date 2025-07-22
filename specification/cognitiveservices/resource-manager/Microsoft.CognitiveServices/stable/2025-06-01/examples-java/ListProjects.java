
/**
 * Samples for Projects List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ListProjects.json
     */
    /**
     * Sample code: List Project.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listProject(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().list("myResourceGroup", "myAccount", com.azure.core.util.Context.NONE);
    }
}
