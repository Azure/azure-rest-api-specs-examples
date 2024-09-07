
/**
 * Samples for PaymentMethods DeleteByUser.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsDeleteByUser.
     * json
     */
    /**
     * Sample code: DeletePaymentMethodOwnedByUser.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void deletePaymentMethodOwnedByUser(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.paymentMethods().deleteByUserWithResponse("ABCDABCDABC0", com.azure.core.util.Context.NONE);
    }
}
