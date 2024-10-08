
/**
 * Samples for PaymentMethods GetByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * paymentMethodsGetByBillingAccount.json
     */
    /**
     * Sample code: PaymentMethodGetAtBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void paymentMethodGetAtBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.paymentMethods().getByBillingAccountWithResponse(
            "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31",
            "21dd9edc-af71-4d62-80ce-37151d475326", com.azure.core.util.Context.NONE);
    }
}
