import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateVCenterRequestProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.VCenter;

/** Samples for ReplicationvCenters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationvCenters_Update.json
     */
    /**
     * Sample code: Update vCenter operation.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updateVCenterOperation(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        VCenter resource =
            manager
                .replicationvCenters()
                .getWithResponse(
                    "MadhaviVault", "MadhaviVRG", "MadhaviFabric", "esx-78", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withProperties(new UpdateVCenterRequestProperties().withIpAddress("10.150.109.25")).apply();
    }
}
