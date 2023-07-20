import com.azure.resourcemanager.managednetworkfabric.models.RebootProperties;
import com.azure.resourcemanager.managednetworkfabric.models.RebootType;

/** Samples for NetworkDevices Reboot. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_Reboot_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_Reboot_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesRebootMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkDevices()
            .reboot(
                "example-rg",
                "example-device",
                new RebootProperties().withRebootType(RebootType.GRACEFUL_REBOOT_WITH_ZTP),
                com.azure.core.util.Context.NONE);
    }
}
