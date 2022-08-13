import com.azure.resourcemanager.automation.models.Sku;
import com.azure.resourcemanager.automation.models.SkuNameEnum;

/** Samples for AutomationAccount CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/createOrUpdateAutomationAccount.json
     */
    /**
     * Sample code: Create or update automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .automationAccounts()
            .define("myAutomationAccount9")
            .withExistingResourceGroup("rg")
            .withRegion("East US 2")
            .withName("myAutomationAccount9")
            .withSku(new Sku().withName(SkuNameEnum.FREE))
            .create();
    }
}
