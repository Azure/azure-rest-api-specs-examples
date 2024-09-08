
/**
 * Samples for PaymentMethods ListByUser.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByUser.json
     */
    /**
     * Sample code: ListPaymentMethodOwnedByUser.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void listPaymentMethodOwnedByUser(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.paymentMethods().listByUser(com.azure.core.util.Context.NONE);
    }
}
