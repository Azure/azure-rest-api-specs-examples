
/**
 * Samples for EmailConfiguration List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/EmailConfiguration_List.json
     */
    /**
     * Sample code: Lists the email configuration settings.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void listsTheEmailConfigurationSettings(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.emailConfigurations().list("rgswagger_2024-09-01", "4", com.azure.core.util.Context.NONE);
    }
}
