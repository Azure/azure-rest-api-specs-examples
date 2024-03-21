
import com.azure.resourcemanager.netapp.models.BreakReplicationRequest;

/**
 * Samples for Volumes BreakReplication.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/Volumes_BreakReplication.json
     */
    /**
     * Sample code: Volumes_BreakReplication.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesBreakReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().breakReplication("myRG", "account1", "pool1", "volume1",
            new BreakReplicationRequest().withForceBreakReplication(false), com.azure.core.util.Context.NONE);
    }
}
