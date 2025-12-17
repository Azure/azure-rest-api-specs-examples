
import com.azure.resourcemanager.networkcloud.models.SkipShutdown;
import com.azure.resourcemanager.networkcloud.models.VirtualMachinePowerOffParameters;

/**
 * Samples for VirtualMachines PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * VirtualMachines_PowerOff.json
     */
    /**
     * Sample code: Power off virtual machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void powerOffVirtualMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().powerOff("resourceGroupName", "virtualMachineName",
            new VirtualMachinePowerOffParameters().withSkipShutdown(SkipShutdown.TRUE),
            com.azure.core.util.Context.NONE);
    }
}
