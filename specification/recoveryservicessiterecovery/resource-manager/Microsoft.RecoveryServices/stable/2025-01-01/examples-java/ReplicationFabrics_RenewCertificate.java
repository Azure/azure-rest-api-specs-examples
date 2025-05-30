
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RenewCertificateInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RenewCertificateInputProperties;

/**
 * Samples for ReplicationFabrics RenewCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationFabrics_RenewCertificate.json
     */
    /**
     * Sample code: Renews certificate for the fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void renewsCertificateForTheFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().renewCertificate("resourceGroupPS1", "vault1", "cloud1",
            new RenewCertificateInput()
                .withProperties(new RenewCertificateInputProperties().withRenewCertificateType("Cloud")),
            com.azure.core.util.Context.NONE);
    }
}
