
/**
 * Samples for BillingAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountsGetEA.json
     */
    /**
     * Sample code: BillingAccountsGetEA.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountsGetEA(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().getWithResponse("6575495", com.azure.core.util.Context.NONE);
    }
}
