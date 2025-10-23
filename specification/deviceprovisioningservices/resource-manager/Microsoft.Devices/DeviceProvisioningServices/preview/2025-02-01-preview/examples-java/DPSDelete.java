
/**
 * Samples for IotDpsResource Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSDelete.json
     */
    /**
     * Sample code: DPSDelete.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSDelete(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().delete("myFirstProvisioningService", "myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
