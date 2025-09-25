
import com.azure.resourcemanager.quota.models.GroupQuotasEntityPatch;
import com.azure.resourcemanager.quota.models.GroupQuotasEntityPatchProperties;

/**
 * Samples for GroupQuotas Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GroupQuotas/PatchGroupQuotas.json
     */
    /**
     * Sample code: GroupQuotas_Patch_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotasPatchRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotas().update("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            new GroupQuotasEntityPatch()
                .withProperties(new GroupQuotasEntityPatchProperties().withDisplayName("UpdatedGroupQuota1")),
            com.azure.core.util.Context.NONE);
    }
}
