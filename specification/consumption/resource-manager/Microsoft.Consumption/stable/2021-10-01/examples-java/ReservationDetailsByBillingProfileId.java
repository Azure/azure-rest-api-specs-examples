import com.azure.core.util.Context;

/** Samples for ReservationsDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetailsByBillingProfileId.json
     */
    /**
     * Sample code: ReservationDetailsByBillingProfileId.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void reservationDetailsByBillingProfileId(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .reservationsDetails()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579",
                "2019-09-01",
                "2019-10-31",
                null,
                null,
                null,
                Context.NONE);
    }
}
