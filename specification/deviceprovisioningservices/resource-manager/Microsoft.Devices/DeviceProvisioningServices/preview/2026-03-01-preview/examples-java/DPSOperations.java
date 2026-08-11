
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DPSOperations.json
     */
    /**
     * Sample code: DPSOperations.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSOperations(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
