/** Samples for ResourceProviderCommon GetSubscriptionQuota. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_usages.json
     */
    /**
     * Sample code: ResourceProviderCommon_GetSubscriptionQuota.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void resourceProviderCommonGetSubscriptionQuota(
        com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.resourceProviderCommons().getSubscriptionQuotaWithResponse(com.azure.core.util.Context.NONE);
    }
}
