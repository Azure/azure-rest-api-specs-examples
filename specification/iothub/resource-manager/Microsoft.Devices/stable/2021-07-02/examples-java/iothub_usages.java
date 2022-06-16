import com.azure.core.util.Context;

/** Samples for ResourceProviderCommon GetSubscriptionQuota. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_usages.json
     */
    /**
     * Sample code: ResourceProviderCommon_GetSubscriptionQuota.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void resourceProviderCommonGetSubscriptionQuota(
        com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.resourceProviderCommons().getSubscriptionQuotaWithResponse(Context.NONE);
    }
}
