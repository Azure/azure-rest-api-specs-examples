
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ConfigureAlertRequestProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationAlertSettings Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationAlertSettings_Create.json
     */
    /**
     * Sample code: Configures email notifications for this vault.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void configuresEmailNotificationsForThisVault(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationAlertSettings().define("defaultAlertSetting").withExistingVault("resourceGroupPS1", "vault1")
            .withProperties(new ConfigureAlertRequestProperties().withSendToOwners("false")
                .withCustomEmailAddresses(Arrays.asList("ronehr@microsoft.com")).withLocale(""))
            .create();
    }
}
