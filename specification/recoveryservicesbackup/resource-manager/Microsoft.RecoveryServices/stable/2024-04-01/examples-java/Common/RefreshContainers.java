
/**
 * Samples for ProtectionContainers Refresh.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * Common/RefreshContainers.json
     */
    /**
     * Sample code: Trigger Azure Vm Discovery.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void triggerAzureVmDiscovery(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionContainers().refreshWithResponse("NetSDKTestRsVault", "SwaggerTestRg", "Azure", null,
            com.azure.core.util.Context.NONE);
    }
}
