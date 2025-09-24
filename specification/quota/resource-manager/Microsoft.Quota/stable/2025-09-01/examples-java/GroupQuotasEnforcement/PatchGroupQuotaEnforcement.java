
import com.azure.resourcemanager.quota.fluent.models.GroupQuotasEnforcementStatusInner;
import com.azure.resourcemanager.quota.models.EnforcementState;
import com.azure.resourcemanager.quota.models.GroupQuotasEnforcementStatusProperties;

/**
 * Samples for GroupQuotaLocationSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GroupQuotasEnforcement/PatchGroupQuotaEnforcement.json
     */
    /**
     * Sample code: GroupQuotaLocationSettings_Patch.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLocationSettingsPatch(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLocationSettings().update("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "eastus",
            new GroupQuotasEnforcementStatusInner().withProperties(
                new GroupQuotasEnforcementStatusProperties().withEnforcementEnabled(EnforcementState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
