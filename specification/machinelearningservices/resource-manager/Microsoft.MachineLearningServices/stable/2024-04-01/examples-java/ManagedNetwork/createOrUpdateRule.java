
import com.azure.resourcemanager.machinelearning.models.FqdnOutboundRule;
import com.azure.resourcemanager.machinelearning.models.RuleCategory;
import com.azure.resourcemanager.machinelearning.models.RuleStatus;

/**
 * Samples for ManagedNetworkSettingsRule CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/ManagedNetwork/createOrUpdateRule.json
     */
    /**
     * Sample code: CreateOrUpdate ManagedNetworkSettingsRule.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateManagedNetworkSettingsRule(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.managedNetworkSettingsRules().define("rule_name_1")
            .withExistingWorkspace("test-rg", "aml-workspace-name")
            .withProperties(new FqdnOutboundRule().withCategory(RuleCategory.USER_DEFINED).withStatus(RuleStatus.ACTIVE)
                .withDestination("destination_endpoint"))
            .create();
    }
}
