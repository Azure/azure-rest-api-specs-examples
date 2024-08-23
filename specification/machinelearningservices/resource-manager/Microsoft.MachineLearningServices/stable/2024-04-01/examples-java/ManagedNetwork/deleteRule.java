
/**
 * Samples for ManagedNetworkSettingsRule Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/ManagedNetwork/deleteRule.json
     */
    /**
     * Sample code: Delete ManagedNetworkSettingsRule.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteManagedNetworkSettingsRule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.managedNetworkSettingsRules().delete("test-rg", "aml-workspace-name", "rule-name",
            com.azure.core.util.Context.NONE);
    }
}
