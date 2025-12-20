
/**
 * Samples for WorkloadNetworks CreatePublicIp.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_CreatePublicIP.json
     */
    /**
     * Sample code: WorkloadNetworks_CreatePublicIP.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreatePublicIP(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().definePublicIp("publicIP1").withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("publicIP1").withNumberOfPublicIPs(32L).create();
    }
}
