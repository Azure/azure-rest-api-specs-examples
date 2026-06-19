
import com.azure.resourcemanager.managednetworkfabric.models.DeviceAdministrativeState;
import com.azure.resourcemanager.managednetworkfabric.models.UpdateDeviceAdministrativeState;
import java.util.Arrays;

/**
 * Samples for NetworkDevices UpdateAdministrativeState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkDevices_UpdateAdministrativeState.json
     */
    /**
     * Sample code: NetworkDevices_UpdateAdministrativeState_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesUpdateAdministrativeStateMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkDevices()
            .updateAdministrativeState("example-rg", "example-device", new UpdateDeviceAdministrativeState()
                .withResourceIds(Arrays.asList("")).withState(DeviceAdministrativeState.RMA),
                com.azure.core.util.Context.NONE);
    }
}
