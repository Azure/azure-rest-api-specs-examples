import com.azure.core.util.Context;

/** Samples for Devices ListRegistrationKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceListRegistrationKey.json
     */
    /**
     * Sample code: Get device registration key.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getDeviceRegistrationKey(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.devices().listRegistrationKeyWithResponse("rg1", "TestDevice", Context.NONE);
    }
}
