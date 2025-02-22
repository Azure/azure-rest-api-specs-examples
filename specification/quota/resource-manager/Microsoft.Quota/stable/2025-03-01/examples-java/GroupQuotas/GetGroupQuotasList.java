
/**
 * Samples for GroupQuotas List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/GroupQuotas/GetGroupQuotasList.
     * json
     */
    /**
     * Sample code: GroupQuotas_List_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotasListRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotas().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", com.azure.core.util.Context.NONE);
    }
}
