
/**
 * Samples for WorkloadNetworks DeletePortMirroring.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_DeletePortMirroring.json
     */
    /**
     * Sample code: WorkloadNetworks_DeletePortMirroring.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeletePortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deletePortMirroring("group1", "portMirroring1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
