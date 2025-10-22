
/**
 * Samples for IotDpsResource List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSListBySubscription.json
     */
    /**
     * Sample code: DPSListBySubscription.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSListBySubscription(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().list(com.azure.core.util.Context.NONE);
    }
}
