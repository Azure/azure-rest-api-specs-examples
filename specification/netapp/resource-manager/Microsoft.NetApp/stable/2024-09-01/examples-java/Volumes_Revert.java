
import com.azure.resourcemanager.netapp.models.VolumeRevert;

/**
 * Samples for Volumes Revert.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/Volumes_Revert.json
     */
    /**
     * Sample code: Volumes_Revert.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesRevert(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().revert("myRG", "account1", "pool1", "volume1", new VolumeRevert().withSnapshotId(
            "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1"),
            com.azure.core.util.Context.NONE);
    }
}
