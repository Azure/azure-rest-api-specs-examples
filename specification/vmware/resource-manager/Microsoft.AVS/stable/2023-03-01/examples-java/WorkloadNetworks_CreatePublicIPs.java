/** Samples for WorkloadNetworks CreatePublicIp. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_CreatePublicIPs.json
     */
    /**
     * Sample code: WorkloadNetworks_CreatePublicIP.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreatePublicIP(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .definePublicIp("publicIP1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("publicIP1")
            .withNumberOfPublicIPs(32L)
            .create();
    }
}
