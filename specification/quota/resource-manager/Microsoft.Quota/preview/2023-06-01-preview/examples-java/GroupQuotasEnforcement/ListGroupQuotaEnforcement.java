
/**
 * Samples for GroupQuotaLocationSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasEnforcement/
     * ListGroupQuotaEnforcement.json
     */
    /**
     * Sample code: GroupQuotaEnforcement_List.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaEnforcementList(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLocationSettings().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", com.azure.core.util.Context.NONE);
    }
}
