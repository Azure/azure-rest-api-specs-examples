
/**
 * Samples for Reservations ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * reservationsListByBillingAccount.json
     */
    /**
     * Sample code: reservationsListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void reservationsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.reservations().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", null, null, null,
            "true", "Succeeded", null, com.azure.core.util.Context.NONE);
    }
}
