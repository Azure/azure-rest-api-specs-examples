
/**
 * Samples for ManagedNetworkSettingsRule Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/ManagedNetwork/getRule.json
     */
    /**
     * Sample code: Get ManagedNetworkSettingsRule.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getManagedNetworkSettingsRule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.managedNetworkSettingsRules().getWithResponse("test-rg", "aml-workspace-name", "name_of_the_fqdn_rule",
            com.azure.core.util.Context.NONE);
    }
}
