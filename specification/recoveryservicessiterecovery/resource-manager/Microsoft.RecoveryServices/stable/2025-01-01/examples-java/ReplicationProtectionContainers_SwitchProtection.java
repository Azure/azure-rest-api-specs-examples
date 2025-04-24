
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2ASwitchProtectionInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.SwitchProtectionInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.SwitchProtectionInputProperties;

/**
 * Samples for ReplicationProtectionContainers SwitchProtection.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionContainers_SwitchProtection.json
     */
    /**
     * Sample code: Switches protection from one container to another or one replication provider to another.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void switchesProtectionFromOneContainerToAnotherOrOneReplicationProviderToAnother(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().switchProtection("priyanprg", "priyanponeboxvault",
            "CentralUSCanSite", "CentralUSCancloud",
            new SwitchProtectionInput()
                .withProperties(new SwitchProtectionInputProperties().withReplicationProtectedItemName("a2aSwapOsVm")
                    .withProviderSpecificDetails(new A2ASwitchProtectionInput())),
            com.azure.core.util.Context.NONE);
    }
}
