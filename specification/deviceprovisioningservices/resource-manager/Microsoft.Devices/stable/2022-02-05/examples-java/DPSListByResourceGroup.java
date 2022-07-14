import com.azure.core.util.Context;

/** Samples for IotDpsResource ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSListByResourceGroup.json
     */
    /**
     * Sample code: DPSListByResourceGroup.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSListByResourceGroup(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
