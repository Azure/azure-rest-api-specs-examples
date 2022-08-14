import com.azure.resourcemanager.automation.fluent.models.RunbookDraftInner;
import com.azure.resourcemanager.automation.models.RunbookTypeEnum;
import java.util.HashMap;
import java.util.Map;

/** Samples for Runbook CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/createRunbookAsDraft.json
     */
    /**
     * Sample code: Create runbook as draft.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createRunbookAsDraft(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .runbooks()
            .define("Get-AzureVMTutorial")
            .withExistingAutomationAccount("rg", "ContoseAutomationAccount")
            .withRunbookType(RunbookTypeEnum.POWER_SHELL_WORKFLOW)
            .withRegion("East US 2")
            .withTags(mapOf("tag01", "value01", "tag02", "value02"))
            .withName("Get-AzureVMTutorial")
            .withLogVerbose(false)
            .withLogProgress(false)
            .withDraft(new RunbookDraftInner())
            .withDescription("Description of the Runbook")
            .create();
    }

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
