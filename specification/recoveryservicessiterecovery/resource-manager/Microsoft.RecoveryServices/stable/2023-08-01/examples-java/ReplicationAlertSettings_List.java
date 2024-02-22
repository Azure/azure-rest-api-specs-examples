
/**
 * Samples for ReplicationAlertSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationAlertSettings_List.json
     */
    /**
     * Sample code: Gets the list of configured email notification(alert) configurations.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfConfiguredEmailNotificationAlertConfigurations(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationAlertSettings().list("vault1", "resourceGroupPS1", com.azure.core.util.Context.NONE);
    }
}
