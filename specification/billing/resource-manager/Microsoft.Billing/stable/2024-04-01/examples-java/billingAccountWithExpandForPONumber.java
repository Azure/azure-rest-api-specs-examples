
/**
 * Samples for BillingAccounts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingAccountWithExpandForPONumber.json
     */
    /**
     * Sample code: BillingAccountWithExpandForPONumber.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountWithExpandForPONumber(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().getWithResponse("8608480", com.azure.core.util.Context.NONE);
    }
}
