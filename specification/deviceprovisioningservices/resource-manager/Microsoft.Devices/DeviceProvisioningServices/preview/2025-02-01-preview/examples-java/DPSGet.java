
/**
 * Samples for IotDpsResource GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGet.json
     */
    /**
     * Sample code: DPSGet.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGet(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().getByResourceGroupWithResponse("myFirstProvisioningService", "myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
