
/**
 * Samples for GroupQuotaUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GroupQuotaUsages/GetGroupQuotaUsages.json
     */
    /**
     * Sample code: GroupQuotasUsages_List.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotasUsagesList(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaUsages().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1", "Microsoft.Compute",
            "westus", com.azure.core.util.Context.NONE);
    }
}
