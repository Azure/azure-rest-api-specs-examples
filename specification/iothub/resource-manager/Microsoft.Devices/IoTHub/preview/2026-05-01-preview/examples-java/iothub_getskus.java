
/**
 * Samples for IotHubResource GetValidSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_getskus.json
     */
    /**
     * Sample code: IotHubResource_GetValidSkus.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetValidSkus(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getValidSkus("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
