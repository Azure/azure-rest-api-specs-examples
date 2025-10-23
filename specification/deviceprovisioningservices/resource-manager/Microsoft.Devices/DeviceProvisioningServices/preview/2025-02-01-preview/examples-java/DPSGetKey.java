
/**
 * Samples for IotDpsResource ListKeysForKeyName.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGetKey.json
     */
    /**
     * Sample code: DPSGetKey.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetKey(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listKeysForKeyNameWithResponse("myFirstProvisioningService", "testKey",
            "myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
