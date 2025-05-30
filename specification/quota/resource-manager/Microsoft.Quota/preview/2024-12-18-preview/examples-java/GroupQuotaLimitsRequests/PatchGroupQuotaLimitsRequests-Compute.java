
import com.azure.resourcemanager.quota.fluent.models.GroupQuotaLimitListInner;
import com.azure.resourcemanager.quota.models.GroupQuotaLimit;
import com.azure.resourcemanager.quota.models.GroupQuotaLimitListProperties;
import com.azure.resourcemanager.quota.models.GroupQuotaLimitProperties;
import java.util.Arrays;

/**
 * Samples for GroupQuotaLimitsRequest Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/GroupQuotaLimitsRequests
     * /PatchGroupQuotaLimitsRequests-Compute.json
     */
    /**
     * Sample code: GroupQuotaLimitsRequests_Update.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaLimitsRequestsUpdate(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaLimitsRequests().update("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "westus",
            new GroupQuotaLimitListInner().withProperties(new GroupQuotaLimitListProperties().withValue(Arrays.asList(
                new GroupQuotaLimit()
                    .withProperties(new GroupQuotaLimitProperties().withResourceName("standardddv4family")
                        .withLimit(110L).withComment("Contoso requires more quota.")),
                new GroupQuotaLimit()
                    .withProperties(new GroupQuotaLimitProperties().withResourceName("standardav2family")
                        .withLimit(110L).withComment("Contoso requires more quota."))))),
            com.azure.core.util.Context.NONE);
    }
}
