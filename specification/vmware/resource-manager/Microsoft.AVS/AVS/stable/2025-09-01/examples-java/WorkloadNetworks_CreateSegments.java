
import com.azure.resourcemanager.avs.models.WorkloadNetworkSegmentSubnet;
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks CreateSegments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_CreateSegments.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateSegments.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateSegments(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().defineSegments("segment1").withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("segment1").withConnectedGateway("/infra/tier-1s/gateway")
            .withSubnet(new WorkloadNetworkSegmentSubnet().withDhcpRanges(Arrays.asList("40.20.0.0-40.20.0.1"))
                .withGatewayAddress("40.20.20.20/16"))
            .withRevision(1L).create();
    }
}
