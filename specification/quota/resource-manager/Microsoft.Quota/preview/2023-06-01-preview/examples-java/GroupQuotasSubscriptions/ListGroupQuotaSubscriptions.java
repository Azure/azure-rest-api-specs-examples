
/**
 * Samples for GroupQuotaSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasSubscriptions
     * /ListGroupQuotaSubscriptions.json
     */
    /**
     * Sample code: GroupQuotaSubscriptions_List_Subscriptions.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaSubscriptionsListSubscriptions(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptions().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            com.azure.core.util.Context.NONE);
    }
}
