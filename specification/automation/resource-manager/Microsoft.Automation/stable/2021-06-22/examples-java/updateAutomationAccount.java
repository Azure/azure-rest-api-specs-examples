
import com.azure.resourcemanager.automation.models.AutomationAccount;
import com.azure.resourcemanager.automation.models.Sku;
import com.azure.resourcemanager.automation.models.SkuNameEnum;

/**
 * Samples for AutomationAccount Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/updateAutomationAccount
     * .json
     */
    /**
     * Sample code: Update an automation account.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void updateAnAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        AutomationAccount resource = manager.automationAccounts()
            .getByResourceGroupWithResponse("rg", "myAutomationAccount9", com.azure.core.util.Context.NONE).getValue();
        resource.update().withName("myAutomationAccount9").withSku(new Sku().withName(SkuNameEnum.FREE)).apply();
    }
}
