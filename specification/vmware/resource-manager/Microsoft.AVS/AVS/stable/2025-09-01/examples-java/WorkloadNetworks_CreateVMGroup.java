
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks CreateVMGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/WorkloadNetworks_CreateVMGroup.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateVMGroup.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateVMGroup(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().defineVMGroup("vmGroup1").withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("vmGroup1").withMembers(Arrays.asList("564d43da-fefc-2a3b-1d92-42855622fa50"))
            .withRevision(1L).create();
    }
}
