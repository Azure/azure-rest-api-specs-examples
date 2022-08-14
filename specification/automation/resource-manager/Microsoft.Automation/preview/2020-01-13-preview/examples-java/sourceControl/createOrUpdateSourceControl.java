import com.azure.resourcemanager.automation.models.SourceControlSecurityTokenProperties;
import com.azure.resourcemanager.automation.models.SourceType;
import com.azure.resourcemanager.automation.models.TokenType;

/** Samples for SourceControl CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/createOrUpdateSourceControl.json
     */
    /**
     * Sample code: Create or update a source control.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateASourceControl(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .sourceControls()
            .define("sampleSourceControl")
            .withExistingAutomationAccount("rg", "sampleAccount9")
            .withRepoUrl("https://sampleUser.visualstudio.com/myProject/_git/myRepository")
            .withBranch("master")
            .withFolderPath("/folderOne/folderTwo")
            .withAutoSync(true)
            .withPublishRunbook(true)
            .withSourceType(SourceType.VSO_GIT)
            .withSecurityToken(
                new SourceControlSecurityTokenProperties()
                    .withAccessToken("3a326f7a0dcd343ea58fee21f2fd5fb4c1234567")
                    .withTokenType(TokenType.PERSONAL_ACCESS_TOKEN))
            .withDescription("my description")
            .create();
    }
}
