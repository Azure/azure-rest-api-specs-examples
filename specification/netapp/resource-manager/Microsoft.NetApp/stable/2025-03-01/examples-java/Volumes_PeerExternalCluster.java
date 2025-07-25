
import com.azure.resourcemanager.netapp.models.PeerClusterForVolumeMigrationRequest;
import java.util.Arrays;

/**
 * Samples for Volumes PeerExternalCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/Volumes_PeerExternalCluster.
     * json
     */
    /**
     * Sample code: Volumes_PeerExternalCluster.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesPeerExternalCluster(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().peerExternalCluster("myRG", "account1", "pool1", "volume1",
            new PeerClusterForVolumeMigrationRequest()
                .withPeerIpAddresses(Arrays.asList("0.0.0.1", "0.0.0.2", "0.0.0.3", "0.0.0.4", "0.0.0.5", "0.0.0.6")),
            com.azure.core.util.Context.NONE);
    }
}
