
import com.azure.resourcemanager.avs.models.WorkloadNetworkSegment;
import com.azure.resourcemanager.avs.models.WorkloadNetworkSegmentSubnet;
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks UpdateSegments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/WorkloadNetworks_UpdateSegments.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdateSegments.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdateSegments(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkSegment resource = manager.workloadNetworks()
            .getSegmentWithResponse("group1", "cloud1", "segment1", com.azure.core.util.Context.NONE).getValue();
        resource
            .update().withConnectedGateway("/infra/tier-1s/gateway").withSubnet(new WorkloadNetworkSegmentSubnet()
                .withDhcpRanges(Arrays.asList("40.20.0.0-40.20.0.1")).withGatewayAddress("40.20.20.20/16"))
            .withRevision(1L).apply();
    }
}
