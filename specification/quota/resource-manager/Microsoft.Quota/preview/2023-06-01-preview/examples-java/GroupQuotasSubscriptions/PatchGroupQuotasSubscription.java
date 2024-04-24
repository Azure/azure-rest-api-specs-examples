
/**
 * Samples for GroupQuotaSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasSubscriptions
     * /PatchGroupQuotasSubscription.json
     */
    /**
     * Sample code: GroupQuotaSubscriptions_Patch_Subscriptions.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaSubscriptionsPatchSubscriptions(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptions().update("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            com.azure.core.util.Context.NONE);
    }
}
