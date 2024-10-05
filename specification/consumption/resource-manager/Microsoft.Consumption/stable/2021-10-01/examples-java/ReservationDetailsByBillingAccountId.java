
/**
 * Samples for ReservationsDetails List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationDetailsByBillingAccountId.json
     */
    /**
     * Sample code: ReservationDetailsByBillingAccountId.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void
        reservationDetailsByBillingAccountId(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationsDetails().list("providers/Microsoft.Billing/billingAccounts/12345", null, null,
            "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05", null, null,
            com.azure.core.util.Context.NONE);
    }
}
