
import com.azure.resourcemanager.managednetworkfabric.models.DeviceRwCommand;

/**
 * Samples for NetworkDevices RunRwCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_RunRwCommand.json
     */
    /**
     * Sample code: NetworkDevices_RunRwCommand_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesRunRwCommandMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices()
            .runRwCommand("example-rg", "example-device", new DeviceRwCommand().withCommand("show running-config")
                .withCommandUrl("https://example.blob.core.windows.net/commands/config.txt"),
                com.azure.core.util.Context.NONE);
    }
}
