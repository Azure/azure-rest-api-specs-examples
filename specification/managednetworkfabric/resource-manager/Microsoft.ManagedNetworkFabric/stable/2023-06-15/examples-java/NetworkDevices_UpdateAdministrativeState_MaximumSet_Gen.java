import com.azure.resourcemanager.managednetworkfabric.models.DeviceAdministrativeState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateDeviceAdministrativeState;
import java.util.Arrays;

/** Samples for NetworkDevices UpdateAdministrativeState. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_UpdateAdministrativeState_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_UpdateAdministrativeState_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkDevices()
            .updateAdministrativeState(
                "example-rg",
                "example-device",
                new UpdateDeviceAdministrativeState()
                    .withResourceIds(Arrays.asList(""))
                    .withState(DeviceAdministrativeState.RMA),
                com.azure.core.util.Context.NONE);
    }
}
