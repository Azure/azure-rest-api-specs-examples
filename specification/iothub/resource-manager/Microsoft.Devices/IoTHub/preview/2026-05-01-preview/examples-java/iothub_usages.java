
/**
 * Samples for ResourceProviderCommon GetSubscriptionQuota.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_usages.json
     */
    /**
     * Sample code: ResourceProviderCommon_GetSubscriptionQuota.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void
        resourceProviderCommonGetSubscriptionQuota(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.resourceProviderCommons().getSubscriptionQuotaWithResponse(com.azure.core.util.Context.NONE);
    }
}
