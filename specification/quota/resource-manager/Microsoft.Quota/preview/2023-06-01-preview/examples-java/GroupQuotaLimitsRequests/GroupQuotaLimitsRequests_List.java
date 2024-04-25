
/**
 * Samples for GroupQuotaLimitsRequest List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests
     * /GroupQuotaLimitsRequests_List.json
     */
    /**
     * Sample code: GroupQuotaLimitsRequest_List.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLimitsRequestList(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLimitsRequests().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "location eq westus", com.azure.core.util.Context.NONE);
    }
}
