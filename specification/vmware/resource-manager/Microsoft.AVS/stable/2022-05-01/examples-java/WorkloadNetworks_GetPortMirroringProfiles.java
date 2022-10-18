import com.azure.core.util.Context;

/** Samples for WorkloadNetworks GetPortMirroring. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetPortMirroringProfiles.json
     */
    /**
     * Sample code: WorkloadNetworks_GetPortMirroring.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetPortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getPortMirroringWithResponse("group1", "cloud1", "portMirroring1", Context.NONE);
    }
}
