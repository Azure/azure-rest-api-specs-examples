import com.azure.core.util.Context;

/** Samples for Devices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceDelete.json
     */
    /**
     * Sample code: Delete device resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteDeviceResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.devices().delete("rg1", "TestDevice", Context.NONE);
    }
}
