
/**
 * Samples for EmailConfiguration Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/EmailConfiguration_Create.json
     */
    /**
     * Sample code: Creates email configuration settings.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void createsEmailConfigurationSettings(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.emailConfigurations().define("0").withExistingReplicationVault("rgswagger_2024-09-01", "4").create();
    }
}
