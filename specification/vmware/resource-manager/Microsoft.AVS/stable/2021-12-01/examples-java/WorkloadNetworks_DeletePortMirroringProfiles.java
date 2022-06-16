import com.azure.core.util.Context;

/** Samples for WorkloadNetworks DeletePortMirroring. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeletePortMirroringProfiles.json
     */
    /**
     * Sample code: WorkloadNetworks_DeletePortMirroring.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeletePortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deletePortMirroring("group1", "portMirroring1", "cloud1", Context.NONE);
    }
}
