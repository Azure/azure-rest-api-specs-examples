
/**
 * Samples for GroupQuotaLimitsRequest Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests
     * /PatchGroupQuotaLimitsRequests-Compute.json
     */
    /**
     * Sample code: GroupQuotaLimitsRequests_Update.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLimitsRequestsUpdate(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLimitsRequests().update("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "standardav2family", null, com.azure.core.util.Context.NONE);
    }
}
