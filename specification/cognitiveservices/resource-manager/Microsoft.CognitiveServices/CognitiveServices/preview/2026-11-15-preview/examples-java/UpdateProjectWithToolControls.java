
import com.azure.resourcemanager.cognitiveservices.models.Project;
import com.azure.resourcemanager.cognitiveservices.models.ProjectProperties;
import com.azure.resourcemanager.cognitiveservices.models.ToolControls;
import com.azure.resourcemanager.cognitiveservices.models.ToolControlsDefaultAction;

/**
 * Samples for Projects Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/UpdateProjectWithToolControls.json
     */
    /**
     * Sample code: Update a Foundry project with no explicitly allowed tools.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateAFoundryProjectWithNoExplicitlyAllowedTools(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Project resource = manager.projects()
            .getWithResponse("myResourceGroup", "myAccount", "myProject", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ProjectProperties()
            .withToolControls(new ToolControls().withDefaultAction(ToolControlsDefaultAction.DENY))).apply();
    }
}
