import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.SourceControl;
import com.azure.resourcemanager.automation.models.SourceControlSecurityTokenProperties;
import com.azure.resourcemanager.automation.models.TokenType;

/** Samples for SourceControl Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/updateSourceControl_patch.json
     */
    /**
     * Sample code: Update a source control.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateASourceControl(com.azure.resourcemanager.automation.AutomationManager manager) {
        SourceControl resource =
            manager
                .sourceControls()
                .getWithResponse("rg", "sampleAccount9", "sampleSourceControl", Context.NONE)
                .getValue();
        resource
            .update()
            .withBranch("master")
            .withFolderPath("/folderOne/folderTwo")
            .withAutoSync(true)
            .withPublishRunbook(true)
            .withSecurityToken(
                new SourceControlSecurityTokenProperties()
                    .withAccessToken("3a326f7a0dcd343ea58fee21f2fd5fb4c1234567")
                    .withTokenType(TokenType.PERSONAL_ACCESS_TOKEN))
            .withDescription("my description")
            .apply();
    }
}
