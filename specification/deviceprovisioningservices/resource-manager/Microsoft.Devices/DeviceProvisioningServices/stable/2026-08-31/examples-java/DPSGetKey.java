
/**
 * Samples for IotDpsResource ListKeysForKeyName.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-31/DPSGetKey.json
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
