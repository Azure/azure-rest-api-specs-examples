
import com.azure.resourcemanager.quota.fluent.models.GroupQuotasEnforcementResponseInner;
import com.azure.resourcemanager.quota.models.EnforcementState;
import com.azure.resourcemanager.quota.models.GroupQuotasEnforcementResponseProperties;

/**
 * Samples for GroupQuotaLocationSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasEnforcement/
     * PutGroupQuotaEnforcementFailed.json
     */
    /**
     * Sample code: GroupQuotaLocationSettings_CreateOrUpdate_Failed.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void
        groupQuotaLocationSettingsCreateOrUpdateFailed(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLocationSettings().createOrUpdate("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "eastus",
            new GroupQuotasEnforcementResponseInner().withProperties(
                new GroupQuotasEnforcementResponseProperties().withEnforcementEnabled(EnforcementState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
