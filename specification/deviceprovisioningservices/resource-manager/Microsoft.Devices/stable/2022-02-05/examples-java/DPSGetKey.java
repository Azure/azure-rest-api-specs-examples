import com.azure.core.util.Context;

/** Samples for IotDpsResource ListKeysForKeyName. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetKey.json
     */
    /**
     * Sample code: DPSGetKey.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetKey(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .listKeysForKeyNameWithResponse("myFirstProvisioningService", "testKey", "myResourceGroup", Context.NONE);
    }
}
