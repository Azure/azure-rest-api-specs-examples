
/**
 * Samples for IotHubResource GetQuotaMetrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_quotametrics.json
     */
    /**
     * Sample code: IotHubResource_GetQuotaMetrics.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetQuotaMetrics(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getQuotaMetrics("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
