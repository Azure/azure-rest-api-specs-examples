import com.azure.core.util.Context;

/** Samples for Devices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/DeviceListByResourceGroup.json
     */
    /**
     * Sample code: List all devices in resource group.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllDevicesInResourceGroup(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.devices().listByResourceGroup("rg1", Context.NONE);
    }
}
