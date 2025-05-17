
import com.azure.resourcemanager.chaos.fluent.models.CapabilityInner;

/**
 * Samples for Capabilities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Capabilities_CreateOrUpdate.json
     */
    /**
     * Sample code: Create/update a Capability that extends a virtual machine Target resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createUpdateACapabilityThatExtendsAVirtualMachineTargetResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.capabilities().createOrUpdateWithResponse("exampleRG", "Microsoft.Compute", "virtualMachines",
            "exampleVM", "Microsoft-VirtualMachine", "Shutdown-1.0", new CapabilityInner(),
            com.azure.core.util.Context.NONE);
    }
}
