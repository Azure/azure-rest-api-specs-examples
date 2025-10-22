
/**
 * Samples for IotDpsResource ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSListKeys.json
     */
    /**
     * Sample code: DPSListKeys.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSListKeys(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listKeys("myFirstProvisioningService", "myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
