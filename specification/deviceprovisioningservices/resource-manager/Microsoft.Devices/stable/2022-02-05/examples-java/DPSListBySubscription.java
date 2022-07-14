import com.azure.core.util.Context;

/** Samples for IotDpsResource List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSListBySubscription.json
     */
    /**
     * Sample code: DPSListBySubscription.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSListBySubscription(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().list(Context.NONE);
    }
}
