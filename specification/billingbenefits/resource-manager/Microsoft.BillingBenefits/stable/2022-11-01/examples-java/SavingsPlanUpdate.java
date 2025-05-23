
import com.azure.resourcemanager.billingbenefits.models.AppliedScopeProperties;
import com.azure.resourcemanager.billingbenefits.models.AppliedScopeType;
import com.azure.resourcemanager.billingbenefits.models.BillingPlan;
import com.azure.resourcemanager.billingbenefits.models.Commitment;
import com.azure.resourcemanager.billingbenefits.models.CommitmentGrain;
import com.azure.resourcemanager.billingbenefits.models.PurchaseRequest;
import com.azure.resourcemanager.billingbenefits.models.RenewProperties;
import com.azure.resourcemanager.billingbenefits.models.SavingsPlanUpdateRequest;
import com.azure.resourcemanager.billingbenefits.models.SavingsPlanUpdateRequestProperties;
import com.azure.resourcemanager.billingbenefits.models.Sku;
import com.azure.resourcemanager.billingbenefits.models.Term;

/**
 * Samples for SavingsPlan Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/
     * SavingsPlanUpdate.json
     */
    /**
     * Sample code: SavingsPlanUpdate.
     * 
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void savingsPlanUpdate(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.savingsPlans().updateWithResponse("20000000-0000-0000-0000-000000000000",
            "30000000-0000-0000-0000-000000000000",
            new SavingsPlanUpdateRequest().withProperties(new SavingsPlanUpdateRequestProperties()
                .withDisplayName("TestDisplayName").withAppliedScopeType(AppliedScopeType.SINGLE)
                .withAppliedScopeProperties(new AppliedScopeProperties()
                    .withResourceGroupId("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"))
                .withRenew(true)
                .withRenewProperties(new RenewProperties().withPurchaseProperties(new PurchaseRequest()
                    .withSku(new Sku().withName("Compute_Savings_Plan")).withDisplayName("TestDisplayName_renewed")
                    .withBillingScopeId("/subscriptions/10000000-0000-0000-0000-000000000000").withTerm(Term.P1Y)
                    .withBillingPlan(BillingPlan.P1M).withAppliedScopeType(AppliedScopeType.SINGLE)
                    .withCommitment(new Commitment()
                        .withCurrencyCode("fakeTokenPlaceholder").withAmount(15.23D).withGrain(CommitmentGrain.HOURLY))
                    .withRenew(false)
                    .withAppliedScopeProperties(new AppliedScopeProperties().withResourceGroupId(
                        "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"))))),
            com.azure.core.util.Context.NONE);
    }
}
