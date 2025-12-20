
/**
 * Samples for WorkloadNetworks ListPortMirroring.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_ListPortMirroring.json
     */
    /**
     * Sample code: WorkloadNetworks_ListPortMirroring.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListPortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listPortMirroring("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
