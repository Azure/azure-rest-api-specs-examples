import com.azure.resourcemanager.avs.models.PortMirroringDirectionEnum;
import com.azure.resourcemanager.avs.models.WorkloadNetworkPortMirroring;

/** Samples for WorkloadNetworks UpdatePortMirroring. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_UpdatePortMirroringProfiles.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdatePortMirroring.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdatePortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkPortMirroring resource =
            manager
                .workloadNetworks()
                .getPortMirroringWithResponse("group1", "cloud1", "portMirroring1", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withDirection(PortMirroringDirectionEnum.BIDIRECTIONAL)
            .withSource("vmGroup1")
            .withDestination("vmGroup2")
            .withRevision(1L)
            .apply();
    }
}
