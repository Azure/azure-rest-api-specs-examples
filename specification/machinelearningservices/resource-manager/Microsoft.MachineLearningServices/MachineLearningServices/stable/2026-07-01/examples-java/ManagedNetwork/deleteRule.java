
/**
 * Samples for ManagedNetworkSettingsRule Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedNetwork/deleteRule.json
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
