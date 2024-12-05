
/**
 * Samples for GroupQuotaSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasSubscriptions
     * /DeleteGroupQuotaSubscriptions.json
     */
    /**
     * Sample code: GroupQuotaSubscriptions_Delete_Subscriptions.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void
        groupQuotaSubscriptionsDeleteSubscriptions(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptions().delete("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            com.azure.core.util.Context.NONE);
    }
}
