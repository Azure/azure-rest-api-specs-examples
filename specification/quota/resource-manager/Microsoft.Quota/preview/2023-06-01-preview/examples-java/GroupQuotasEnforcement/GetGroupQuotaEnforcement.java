
/**
 * Samples for GroupQuotaLocationSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasEnforcement/
     * GetGroupQuotaEnforcement.json
     */
    /**
     * Sample code: GroupQuotasEnforcement_Get.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotasEnforcementGet(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLocationSettings().getWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "eastus", com.azure.core.util.Context.NONE);
    }
}
