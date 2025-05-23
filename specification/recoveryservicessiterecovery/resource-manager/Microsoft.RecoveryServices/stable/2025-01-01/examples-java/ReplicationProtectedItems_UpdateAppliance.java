
import com.azure.resourcemanager.recoveryservicessiterecovery.models.InMageRcmUpdateApplianceForReplicationProtectedItemInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateApplianceForReplicationProtectedItemInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateApplianceForReplicationProtectedItemInputProperties;

/**
 * Samples for ReplicationProtectedItems UpdateAppliance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_UpdateAppliance.json
     */
    /**
     * Sample code: Updates appliance for replication protected Item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updatesApplianceForReplicationProtectedItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().updateAppliance("Ayan-0106-SA-RG", "Ayan-0106-SA-Vault",
            "Ayan-0106-SA-Vaultreplicationfabric", "Ayan-0106-SA-Vaultreplicationcontainer",
            "idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f",
            new UpdateApplianceForReplicationProtectedItemInput()
                .withProperties(new UpdateApplianceForReplicationProtectedItemInputProperties()
                    .withTargetApplianceId("").withProviderSpecificDetails(
                        new InMageRcmUpdateApplianceForReplicationProtectedItemInput().withRunAsAccountId(""))),
            com.azure.core.util.Context.NONE);
    }
}
