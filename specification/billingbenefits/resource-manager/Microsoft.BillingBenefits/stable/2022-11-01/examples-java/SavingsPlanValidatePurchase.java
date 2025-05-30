
import com.azure.resourcemanager.billingbenefits.fluent.models.SavingsPlanOrderAliasModelInner;
import com.azure.resourcemanager.billingbenefits.models.AppliedScopeProperties;
import com.azure.resourcemanager.billingbenefits.models.AppliedScopeType;
import com.azure.resourcemanager.billingbenefits.models.Commitment;
import com.azure.resourcemanager.billingbenefits.models.CommitmentGrain;
import com.azure.resourcemanager.billingbenefits.models.SavingsPlanPurchaseValidateRequest;
import com.azure.resourcemanager.billingbenefits.models.Sku;
import com.azure.resourcemanager.billingbenefits.models.Term;
import java.util.Arrays;

/**
 * Samples for ResourceProvider ValidatePurchase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/
     * SavingsPlanValidatePurchase.json
     */
    /**
     * Sample code: SavingsPlanValidatePurchase.
     * 
     * @param manager Entry point to BillingBenefitsManager.
     */
    public static void
        savingsPlanValidatePurchase(com.azure.resourcemanager.billingbenefits.BillingBenefitsManager manager) {
        manager.resourceProviders()
            .validatePurchaseWithResponse(new SavingsPlanPurchaseValidateRequest().withBenefits(Arrays.asList(
                new SavingsPlanOrderAliasModelInner().withSku(new Sku().withName("Compute_Savings_Plan"))
                    .withDisplayName("ComputeSavingsPlan_2021-07-01")
                    .withBillingScopeId("/subscriptions/10000000-0000-0000-0000-000000000000").withTerm(Term.P1Y)
                    .withAppliedScopeType(AppliedScopeType.SINGLE)
                    .withAppliedScopeProperties(new AppliedScopeProperties().withResourceGroupId(
                        "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"))
                    .withCommitment(new Commitment().withCurrencyCode("fakeTokenPlaceholder").withAmount(15.23D)
                        .withGrain(CommitmentGrain.HOURLY)),
                new SavingsPlanOrderAliasModelInner().withSku(new Sku().withName("Compute_Savings_Plan"))
                    .withDisplayName("ComputeSavingsPlan_2021-07-01")
                    .withBillingScopeId("/subscriptions/10000000-0000-0000-0000-000000000000").withTerm(Term.P1Y)
                    .withAppliedScopeType(AppliedScopeType.SINGLE)
                    .withAppliedScopeProperties(new AppliedScopeProperties()
                        .withResourceGroupId("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/RG"))
                    .withCommitment(new Commitment().withCurrencyCode("fakeTokenPlaceholder").withAmount(20.0D)
                        .withGrain(CommitmentGrain.HOURLY)))),
                com.azure.core.util.Context.NONE);
    }
}
