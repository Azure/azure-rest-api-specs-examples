
/**
 * Samples for GroupQuotaLimits Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimits/
     * GetGroupQuotaLimits-Compute.json
     */
    /**
     * Sample code: GroupQuotaLimits_Get_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLimitsGetRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLimits().getWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "cores", "location eq westus", com.azure.core.util.Context.NONE);
    }
}
