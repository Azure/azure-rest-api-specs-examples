
import com.azure.resourcemanager.networkcloud.models.BareMetalMachinePowerOffParameters;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineSkipShutdown;

/**
 * Samples for BareMetalMachines PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachines_PowerOff.json
     */
    /**
     * Sample code: Power off bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void powerOffBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().powerOff("resourceGroupName", "bareMetalMachineName",
            new BareMetalMachinePowerOffParameters().withSkipShutdown(BareMetalMachineSkipShutdown.TRUE),
            com.azure.core.util.Context.NONE);
    }
}
