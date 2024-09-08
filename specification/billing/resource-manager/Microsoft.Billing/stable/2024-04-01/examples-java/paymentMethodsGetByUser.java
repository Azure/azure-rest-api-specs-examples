
/**
 * Samples for PaymentMethods GetByUser.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByUser.json
     */
    /**
     * Sample code: GetPaymentMethodOwnedByUser.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void getPaymentMethodOwnedByUser(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.paymentMethods().getByUserWithResponse("ABCDABCDABC0", com.azure.core.util.Context.NONE);
    }
}
