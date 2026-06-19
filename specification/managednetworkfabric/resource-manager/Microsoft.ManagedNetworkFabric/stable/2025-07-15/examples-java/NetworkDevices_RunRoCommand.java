
import com.azure.resourcemanager.managednetworkfabric.models.DeviceRoCommand;

/**
 * Samples for NetworkDevices RunRoCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_RunRoCommand.json
     */
    /**
     * Sample code: NetworkDevices_RunRoCommand_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesRunRoCommandMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices().runRoCommand("example-rg", "example-device",
            new DeviceRoCommand().withCommand("show version"), com.azure.core.util.Context.NONE);
    }
}
