
import com.azure.resourcemanager.automation.models.ContentHash;
import com.azure.resourcemanager.automation.models.ContentLink;
import com.azure.resourcemanager.automation.models.RunbookTypeEnum;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Runbook CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createOrUpdateRunbook.
     * json
     */
    /**
     * Sample code: Create or update runbook and publish it.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void
        createOrUpdateRunbookAndPublishIt(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbooks().define("Get-AzureVMTutorial").withExistingAutomationAccount("rg", "ContoseAutomationAccount")
            .withRunbookType(RunbookTypeEnum.POWER_SHELL_WORKFLOW).withRegion("East US 2")
            .withTags(mapOf("tag01", "value01", "tag02", "value02")).withName("Get-AzureVMTutorial")
            .withLogVerbose(false).withLogProgress(true)
            .withPublishContentLink(new ContentLink().withUri(
                "https://raw.githubusercontent.com/Azure/azure-quickstart-templates/master/101-automation-runbook-getvms/Runbooks/Get-AzureVMTutorial.ps1")
                .withContentHash(new ContentHash().withAlgorithm("SHA256")
                    .withValue("115775B8FF2BE672D8A946BD0B489918C724DDE15A440373CA54461D53010A80")))
            .withDescription("Description of the Runbook").withLogActivityTrace(1).create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
