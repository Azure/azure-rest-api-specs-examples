
import com.azure.resourcemanager.avs.models.VirtualMachineRestrictMovement;
import com.azure.resourcemanager.avs.models.VirtualMachineRestrictMovementState;

/**
 * Samples for VirtualMachines RestrictMovement.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/VirtualMachines_RestrictMovement.
     * json
     */
    /**
     * Sample code: VirtualMachines_RestrictMovement.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void virtualMachinesRestrictMovement(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.virtualMachines().restrictMovement("group1", "cloud1", "cluster1", "vm-209",
            new VirtualMachineRestrictMovement().withRestrictMovement(VirtualMachineRestrictMovementState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
