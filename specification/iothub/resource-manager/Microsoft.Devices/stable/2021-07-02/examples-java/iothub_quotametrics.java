import com.azure.core.util.Context;

/** Samples for IotHubResource GetQuotaMetrics. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_quotametrics.json
     */
    /**
     * Sample code: IotHubResource_GetQuotaMetrics.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetQuotaMetrics(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getQuotaMetrics("myResourceGroup", "testHub", Context.NONE);
    }
}
