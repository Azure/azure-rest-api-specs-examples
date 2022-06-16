import com.azure.core.util.Context;

/** Samples for IotDpsResource ListValidSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSGetValidSku.json
     */
    /**
     * Sample code: DPSGetValidSku.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetValidSku(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listValidSkus("myFirstProvisioningService", "myResourceGroup", Context.NONE);
    }
}
