
/**
 * Samples for IotDpsResource GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DPSGet.json
     */
    /**
     * Sample code: DPSGet.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGet(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().getByResourceGroupWithResponse("myResourceGroup", "myFirstProvisioningService",
            com.azure.core.util.Context.NONE);
    }
}
