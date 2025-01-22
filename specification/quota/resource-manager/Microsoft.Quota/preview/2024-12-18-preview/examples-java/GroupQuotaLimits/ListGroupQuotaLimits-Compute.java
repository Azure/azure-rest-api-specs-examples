
/**
 * Samples for GroupQuotaLimits List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/GroupQuotaLimits/
     * ListGroupQuotaLimits-Compute.json
     */
    /**
     * Sample code: GroupQuotaLimits_Get_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLimitsGetRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLimits().listWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "westus", com.azure.core.util.Context.NONE);
    }
}
