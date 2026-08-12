
/**
 * Samples for IotDpsResource ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/DPSListByResourceGroup.json
     */
    /**
     * Sample code: DPSListByResourceGroup.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSListByResourceGroup(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
