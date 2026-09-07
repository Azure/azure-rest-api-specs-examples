
/**
 * Samples for IotDpsResource GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-31/DPSGet_DisableLocalAuth.json
     */
    /**
     * Sample code: DPSGet_DisableLocalAuth.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSGetDisableLocalAuth(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().getByResourceGroupWithResponse("myResourceGroup", "myFirstProvisioningService",
            com.azure.core.util.Context.NONE);
    }
}
