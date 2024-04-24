
/**
 * Samples for GroupQuotaSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasSubscriptions
     * /PutGroupQuotasSubscription.json
     */
    /**
     * Sample code: GroupQuotaSubscriptions_Put_Subscriptions.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaSubscriptionsPutSubscriptions(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptions().createOrUpdate("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            com.azure.core.util.Context.NONE);
    }
}
