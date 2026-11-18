
import com.azure.resourcemanager.cognitiveservices.models.ProjectProperties;
import com.azure.resourcemanager.cognitiveservices.models.ToolControlEntry;
import com.azure.resourcemanager.cognitiveservices.models.ToolControlType;
import com.azure.resourcemanager.cognitiveservices.models.ToolControls;
import com.azure.resourcemanager.cognitiveservices.models.ToolControlsDefaultAction;
import java.util.Arrays;

/**
 * Samples for Projects Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/CreateProjectWithToolControls.json
     */
    /**
     * Sample code: Create a Foundry project with tool controls.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAFoundryProjectWithToolControls(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().define("myProject").withExistingAccount("myResourceGroup", "myAccount").withRegion("westus")
            .withProperties(new ProjectProperties().withDisplayName("My project").withToolControls(new ToolControls()
                .withDefaultAction(ToolControlsDefaultAction.DENY)
                .withAllowedTools(Arrays.asList(new ToolControlEntry().withType(ToolControlType.CODE_INTERPRETER),
                    new ToolControlEntry().withType(ToolControlType.MCP).withServerUrl("https://*.contoso.com/mcp")))))
            .create();
    }
}
