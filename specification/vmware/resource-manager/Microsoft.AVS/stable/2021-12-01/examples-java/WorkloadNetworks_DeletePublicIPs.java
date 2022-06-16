import com.azure.core.util.Context;

/** Samples for WorkloadNetworks DeletePublicIp. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeletePublicIPs.json
     */
    /**
     * Sample code: WorkloadNetworks_DeletePublicIP.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksDeletePublicIP(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().deletePublicIp("group1", "publicIP1", "cloud1", Context.NONE);
    }
}
