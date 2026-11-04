
import com.azure.resourcemanager.recoveryservicessiterecovery.models.DisableProtectionInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.DisableProtectionInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.DisableProtectionProviderSpecificInput;

/**
 * Samples for ReplicationProtectedItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Delete.json
     */
    /**
     * Sample code: Disables protection.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        disablesProtection(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().delete("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "c0c14913-3d7a-48ea-9531-cc99e0e686e6",
            new DisableProtectionInput().withProperties(new DisableProtectionInputProperties()
                .withReplicationProviderInput(new DisableProtectionProviderSpecificInput())),
            com.azure.core.util.Context.NONE);
    }
}
