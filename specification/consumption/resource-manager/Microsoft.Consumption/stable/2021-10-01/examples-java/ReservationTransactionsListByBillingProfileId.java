
/**
 * Samples for ReservationTransactions ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ReservationTransactionsListByBillingProfileId.json
     */
    /**
     * Sample code: ReservationTransactionsByBillingProfileId.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void
        reservationTransactionsByBillingProfileId(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.reservationTransactions().listByBillingProfile(
            "fcebaabc-fced-4284-a83d-79f83dee183c:45796ba8-988f-45ad-bea9-7b71fc6c7513_2018-09-30", "Z76D-SGAF-BG7-TGB",
            "properties/eventDate ge 2020-05-20 AND properties/eventDate le 2020-05-30",
            com.azure.core.util.Context.NONE);
    }
}
