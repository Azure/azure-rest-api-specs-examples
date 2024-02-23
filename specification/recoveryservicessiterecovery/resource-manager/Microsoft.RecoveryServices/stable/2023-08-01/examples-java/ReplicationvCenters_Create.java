
import com.azure.resourcemanager.recoveryservicessiterecovery.models.AddVCenterRequestProperties;

/**
 * Samples for ReplicationvCenters Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationvCenters_Create.json
     */
    /**
     * Sample code: Add vCenter.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void addVCenter(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationvCenters().define("esx-78")
            .withExistingReplicationFabric("MadhaviVault", "MadhaviVRG", "MadhaviFabric")
            .withProperties(new AddVCenterRequestProperties().withFriendlyName("esx-78").withIpAddress("inmtest78")
                .withProcessServerId("5A720CAB-39CB-F445-BD1662B0B33164B5").withPort("443").withRunAsAccountId("2"))
            .create();
    }
}
