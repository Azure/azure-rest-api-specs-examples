
import com.azure.resourcemanager.avs.models.WorkloadNetworkVMGroup;
import java.util.Arrays;

/**
 * Samples for WorkloadNetworks UpdateVMGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/WorkloadNetworks_UpdateVMGroup.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdateVMGroup.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdateVMGroup(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkVMGroup resource = manager.workloadNetworks()
            .getVMGroupWithResponse("group1", "cloud1", "vmGroup1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withMembers(Arrays.asList("564d43da-fefc-2a3b-1d92-42855622fa50")).withRevision(1L).apply();
    }
}
