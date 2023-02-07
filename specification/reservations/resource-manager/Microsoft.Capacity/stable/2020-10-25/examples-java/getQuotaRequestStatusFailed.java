/** Samples for QuotaRequestStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getQuotaRequestStatusFailed.json
     */
    /**
     * Sample code: QuotaRequestFailed.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotaRequestFailed(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .quotaRequestStatus()
            .getWithResponse(
                "00000000-0000-0000-0000-000000000000",
                "Microsoft.Compute",
                "eastus",
                "2B5C8515-37D8-4B6A-879B-CD641A2CF605",
                com.azure.core.util.Context.NONE);
    }
}
