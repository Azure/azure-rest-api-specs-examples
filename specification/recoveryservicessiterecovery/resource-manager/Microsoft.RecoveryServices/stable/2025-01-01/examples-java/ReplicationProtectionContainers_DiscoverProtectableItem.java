
import com.azure.resourcemanager.recoveryservicessiterecovery.models.DiscoverProtectableItemRequest;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.DiscoverProtectableItemRequestProperties;

/**
 * Samples for ReplicationProtectionContainers DiscoverProtectableItem.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionContainers_DiscoverProtectableItem.json
     */
    /**
     * Sample code: Adds a protectable item to the replication protection container.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void addsAProtectableItemToTheReplicationProtectionContainer(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().discoverProtectableItem("MadhaviVRG", "MadhaviVault", "V2A-W2K12-660",
            "cloud_7328549c-5c37-4459-a3c2-e35f9ef6893c",
            new DiscoverProtectableItemRequest().withProperties(new DiscoverProtectableItemRequestProperties()
                .withFriendlyName("Test").withIpAddress("10.150.2.3").withOsType("Windows")),
            com.azure.core.util.Context.NONE);
    }
}
