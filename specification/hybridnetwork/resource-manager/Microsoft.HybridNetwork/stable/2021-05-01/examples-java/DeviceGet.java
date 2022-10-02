import com.azure.core.util.Context;

/** Samples for Devices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceGet.json
     */
    /**
     * Sample code: Get device resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getDeviceResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.devices().getByResourceGroupWithResponse("rg1", "TestDevice", Context.NONE);
    }
}
