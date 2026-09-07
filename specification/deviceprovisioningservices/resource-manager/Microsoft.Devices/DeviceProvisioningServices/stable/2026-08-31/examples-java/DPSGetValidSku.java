
/**
 * Samples for IotDpsResource ListValidSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-31/DPSGetValidSku.json
     */
    /**
     * Sample code: DPSGetValidSku.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetValidSku(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listValidSkus("myFirstProvisioningService", "myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
