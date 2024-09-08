
/**
 * Samples for BillingProfiles ValidateDeleteEligibility.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingProfilesValidateDeleteEligibilitySuccess.json
     */
    /**
     * Sample code: BillingProfilesValidateDeleteEligibilitySuccess.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingProfilesValidateDeleteEligibilitySuccess(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProfiles().validateDeleteEligibilityWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            com.azure.core.util.Context.NONE);
    }
}
