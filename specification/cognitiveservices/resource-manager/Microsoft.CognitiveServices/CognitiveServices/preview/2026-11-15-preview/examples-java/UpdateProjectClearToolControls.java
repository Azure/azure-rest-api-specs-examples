
import com.azure.resourcemanager.cognitiveservices.models.Project;
import com.azure.resourcemanager.cognitiveservices.models.ProjectProperties;

/**
 * Samples for Projects Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/UpdateProjectClearToolControls.json
     */
    /**
     * Sample code: Clear a Foundry project's tool controls with explicit null.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void clearAFoundryProjectSToolControlsWithExplicitNull(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Project resource = manager.projects()
            .getWithResponse("myResourceGroup", "myAccount", "myProject", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ProjectProperties()).apply();
    }
}
