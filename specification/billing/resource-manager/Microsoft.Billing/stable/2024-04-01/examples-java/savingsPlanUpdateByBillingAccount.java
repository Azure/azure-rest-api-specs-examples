
import com.azure.resourcemanager.billing.models.AppliedScopeProperties;
import com.azure.resourcemanager.billing.models.AppliedScopeType;
import com.azure.resourcemanager.billing.models.BillingPlan;
import com.azure.resourcemanager.billing.models.Commitment;
import com.azure.resourcemanager.billing.models.CommitmentGrain;
import com.azure.resourcemanager.billing.models.PurchaseRequest;
import com.azure.resourcemanager.billing.models.RenewProperties;
import com.azure.resourcemanager.billing.models.SavingsPlanTerm;
import com.azure.resourcemanager.billing.models.SavingsPlanUpdateRequest;
import com.azure.resourcemanager.billing.models.SavingsPlanUpdateRequestProperties;
import com.azure.resourcemanager.billing.models.Sku;

/**
 * Samples for SavingsPlans UpdateByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * savingsPlanUpdateByBillingAccount.json
     */
    /**
     * Sample code: SavingsPlanUpdate.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void savingsPlanUpdate(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.savingsPlans()
            .updateByBillingAccount(
                "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
                "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000",
                new SavingsPlanUpdateRequest()
                    .withProperties(
                        new SavingsPlanUpdateRequestProperties().withDisplayName("sp_newName")
                            .withAppliedScopeType(AppliedScopeType.MANAGEMENT_GROUP)
                            .withAppliedScopeProperties(new AppliedScopeProperties()
                                .withTenantId("80000000-0000-0000-0000-000000000000")
                                .withManagementGroupId("/providers/Microsoft.Management/managementGroups/mg1"))
                            .withRenew(true)
                            .withRenewProperties(new RenewProperties().withPurchaseProperties(new PurchaseRequest()
                                .withSku(new Sku().withName("Compute_Savings_Plan"))
                                .withDisplayName("sp_newName_renewed")
                                .withBillingScopeId("/subscriptions/50000000-0000-0000-0000-000000000000")
                                .withTerm(SavingsPlanTerm.P3Y).withBillingPlan(BillingPlan.P1M)
                                .withAppliedScopeType(AppliedScopeType.MANAGEMENT_GROUP)
                                .withCommitment(new Commitment().withCurrencyCode("fakeTokenPlaceholder")
                                    .withAmount(0.001D).withGrain(CommitmentGrain.HOURLY))
                                .withAppliedScopeProperties(new AppliedScopeProperties()
                                    .withTenantId("80000000-0000-0000-0000-000000000000")
                                    .withManagementGroupId("/providers/Microsoft.Management/managementGroups/mg1"))))),
                com.azure.core.util.Context.NONE);
    }
}
