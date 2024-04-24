
/**
 * Samples for GroupQuotas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotas/
     * GetGroupQuotas.json
     */
    /**
     * Sample code: GroupQuotas_Get_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotasGetRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotas().getWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            com.azure.core.util.Context.NONE);
    }
}
