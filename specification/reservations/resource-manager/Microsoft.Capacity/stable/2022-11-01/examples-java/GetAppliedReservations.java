/** Samples for ResourceProvider GetAppliedReservationList. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetAppliedReservations.json
     */
    /**
     * Sample code: AppliedReservationList.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void appliedReservationList(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .resourceProviders()
            .getAppliedReservationListWithResponse(
                "23bc208b-083f-4901-ae85-4f98c0c3b4b6", com.azure.core.util.Context.NONE);
    }
}
