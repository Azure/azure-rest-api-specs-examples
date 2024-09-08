
/**
 * Samples for PaymentMethods ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * paymentMethodsListByBillingAccount.json
     */
    /**
     * Sample code: PaymentMethodsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void paymentMethodsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.paymentMethods().listByBillingAccount(
            "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31",
            com.azure.core.util.Context.NONE);
    }
}
