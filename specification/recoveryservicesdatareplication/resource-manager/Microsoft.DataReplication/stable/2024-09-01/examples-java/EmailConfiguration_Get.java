
/**
 * Samples for EmailConfiguration Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/EmailConfiguration_Get.json
     */
    /**
     * Sample code: Gets the email configuration setting.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getsTheEmailConfigurationSetting(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.emailConfigurations().getWithResponse("rgswagger_2024-09-01", "4", "0",
            com.azure.core.util.Context.NONE);
    }
}
