
/**
 * Samples for GroupQuotaLimitsRequest Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests
     * /GroupQuotaLimitsRequests_Get.json
     */
    /**
     * Sample code: GroupQuotaLimitsRequests_Get.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLimitsRequestsGet(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLimitsRequests().getWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "requestId", com.azure.core.util.Context.NONE);
    }
}
