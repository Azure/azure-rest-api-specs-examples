import com.azure.core.util.Context;

/** Samples for WorkloadNetworks ListPortMirroring. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListPortMirroringProfiles.json
     */
    /**
     * Sample code: WorkloadNetworks_ListPortMirroring.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListPortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listPortMirroring("group1", "cloud1", Context.NONE);
    }
}
