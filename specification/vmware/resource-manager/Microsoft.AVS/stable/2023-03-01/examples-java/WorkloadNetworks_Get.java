import com.azure.resourcemanager.avs.models.WorkloadNetworkName;

/** Samples for WorkloadNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_Get.json
     */
    /**
     * Sample code: WorkloadNetworks_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .getWithResponse("group1", "cloud1", WorkloadNetworkName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
