import com.azure.resourcemanager.recoveryservicessiterecovery.models.ConfigureAlertRequestProperties;
import java.util.Arrays;

/** Samples for ReplicationAlertSettings Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationAlertSettings_Create.json
     */
    /**
     * Sample code: Configures email notifications for this vault.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void configuresEmailNotificationsForThisVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationAlertSettings()
            .define("defaultAlertSetting")
            .withExistingVault("vault1", "resourceGroupPS1")
            .withProperties(
                new ConfigureAlertRequestProperties()
                    .withSendToOwners("false")
                    .withCustomEmailAddresses(Arrays.asList("ronehr@microsoft.com"))
                    .withLocale(""))
            .create();
    }
}
