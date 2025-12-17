
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineCordonParameters;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineEvacuate;

/**
 * Samples for BareMetalMachines Cordon.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachines_Cordon.json
     */
    /**
     * Sample code: Cordon bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void cordonBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().cordon("resourceGroupName", "bareMetalMachineName",
            new BareMetalMachineCordonParameters().withEvacuate(BareMetalMachineEvacuate.TRUE),
            com.azure.core.util.Context.NONE);
    }
}
