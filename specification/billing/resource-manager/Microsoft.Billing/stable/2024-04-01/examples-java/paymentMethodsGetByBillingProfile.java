
/**
 * Samples for PaymentMethods GetByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * paymentMethodsGetByBillingProfile.json
     */
    /**
     * Sample code: PaymentMethodsGetByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void paymentMethodsGetByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.paymentMethods().getByBillingProfileWithResponse(
            "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31", "ABC1-A1CD-AB1-BP1",
            "ABCDABCDABC0", com.azure.core.util.Context.NONE);
    }
}
