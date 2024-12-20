
import com.azure.resourcemanager.billingbenefits.models.AppliedScopeProperties;
import com.azure.resourcemanager.billingbenefits.models.AppliedScopeType;
import com.azure.resourcemanager.billingbenefits.models.SavingsPlanUpdateRequestProperties;
import com.azure.resourcemanager.billingbenefits.models.SavingsPlanUpdateValidateRequest;
import java.util.Arrays;

/**
 * Samples for SavingsPlan ValidateUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/
     * SavingsPlanValidateUpdate.json
     */
    /**
     * Sample code: SavingsPlanValidateUpdate.
     * 
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void
        savingsPlanValidateUpdate(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.savingsPlans().validateUpdateWithResponse("20000000-0000-0000-0000-000000000000",
            "30000000-0000-0000-0000-000000000000",
            new SavingsPlanUpdateValidateRequest().withBenefits(Arrays.asList(
                new SavingsPlanUpdateRequestProperties().withAppliedScopeType(AppliedScopeType.MANAGEMENT_GROUP)
                    .withAppliedScopeProperties(new AppliedScopeProperties()
                        .withTenantId("30000000-0000-0000-0000-000000000100").withManagementGroupId(
                            "/providers/Microsoft.Management/managementGroups/30000000-0000-0000-0000-000000000100")),
                new SavingsPlanUpdateRequestProperties().withAppliedScopeType(AppliedScopeType.MANAGEMENT_GROUP)
                    .withAppliedScopeProperties(
                        new AppliedScopeProperties().withTenantId("30000000-0000-0000-0000-000000000100")
                            .withManagementGroupId("/providers/Microsoft.Management/managementGroups/MockMG")))),
            com.azure.core.util.Context.NONE);
    }
}
