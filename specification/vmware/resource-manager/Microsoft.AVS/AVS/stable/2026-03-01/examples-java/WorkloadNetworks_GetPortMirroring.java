
/**
 * Samples for WorkloadNetworks GetPortMirroring.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/WorkloadNetworks_GetPortMirroring.json
     */
    /**
     * Sample code: WorkloadNetworks_GetPortMirroring.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetPortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getPortMirroringWithResponse("group1", "cloud1", "portMirroring1",
            com.azure.core.util.Context.NONE);
    }
}
