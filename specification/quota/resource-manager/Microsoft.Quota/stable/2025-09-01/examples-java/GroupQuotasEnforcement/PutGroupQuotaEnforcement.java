
import com.azure.resourcemanager.quota.fluent.models.GroupQuotasEnforcementStatusInner;
import com.azure.resourcemanager.quota.models.EnforcementState;
import com.azure.resourcemanager.quota.models.GroupQuotasEnforcementStatusProperties;

/**
 * Samples for GroupQuotaLocationSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GroupQuotasEnforcement/PutGroupQuotaEnforcement.json
     */
    /**
     * Sample code: GroupQuotaLocationSettings_CreateOrUpdate.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLocationSettingsCreateOrUpdate(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLocationSettings().createOrUpdate("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "eastus",
            new GroupQuotasEnforcementStatusInner().withProperties(
                new GroupQuotasEnforcementStatusProperties().withEnforcementEnabled(EnforcementState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
