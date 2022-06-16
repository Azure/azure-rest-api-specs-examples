import com.azure.core.util.Context;

/** Samples for IotDpsResource Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSDelete.json
     */
    /**
     * Sample code: DPSDelete.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSDelete(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().delete("myResourceGroup", "myFirstProvisioningService", Context.NONE);
    }
}
