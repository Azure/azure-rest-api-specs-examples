
/**
 * Samples for Reservations ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * reservationsListByBillingProfile.json
     */
    /**
     * Sample code: ReservationsListByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void reservationsListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.reservations().listByBillingProfile(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "AAAA-AAAA-AAA-AAA",
            null, null, null, "true", "Succeeded", null, com.azure.core.util.Context.NONE);
    }
}
