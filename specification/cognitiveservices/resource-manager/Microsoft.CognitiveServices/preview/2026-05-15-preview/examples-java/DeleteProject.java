
/**
 * Samples for Projects Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/DeleteProject.json
     */
    /**
     * Sample code: Delete Project.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteProject(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().delete("myResourceGroup", "PropTest01", "myProject", com.azure.core.util.Context.NONE);
    }
}
