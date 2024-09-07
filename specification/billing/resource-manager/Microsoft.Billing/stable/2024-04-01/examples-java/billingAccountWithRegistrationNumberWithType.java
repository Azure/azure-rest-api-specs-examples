
/**
 * Samples for BillingAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingAccountWithRegistrationNumberWithType.json
     */
    /**
     * Sample code: BillingAccountWithRegistrationNumberWithType.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingAccountWithRegistrationNumberWithType(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().getWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            com.azure.core.util.Context.NONE);
    }
}
