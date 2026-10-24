
/**
 * Samples for ProtectionContainers Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Common/RefreshContainers_CrossSubscription.json
     */
    /**
     * Sample code: Trigger cross-subscription Azure Workload discovery.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void triggerCrossSubscriptionAzureWorkloadDiscovery(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionContainers().refreshWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "Azure",
            "backupManagementType eq 'AzureWorkload' and containerSubscriptionId eq 'a76f4f58-8c04-4f53-9e68-4a698b0f43e4'",
            com.azure.core.util.Context.NONE);
    }
}
