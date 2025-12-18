
import com.azure.resourcemanager.networkcloud.models.RelayType;
import com.azure.resourcemanager.networkcloud.models.VirtualMachineAssignRelayParameters;

/**
 * Samples for VirtualMachines AssignRelay.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * VirtualMachines_AssignRelay.json
     */
    /**
     * Sample code: Assign relay to the Microsoft.HybridCompute machine for a virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void assignRelayToTheMicrosoftHybridComputeMachineForAVirtualMachine(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().assignRelay("resourceGroupName", "virtualMachineName",
            new VirtualMachineAssignRelayParameters().withMachineId(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.HybridCompute/machines/machineName")
                .withRelayType(RelayType.PLATFORM),
            com.azure.core.util.Context.NONE);
    }
}
