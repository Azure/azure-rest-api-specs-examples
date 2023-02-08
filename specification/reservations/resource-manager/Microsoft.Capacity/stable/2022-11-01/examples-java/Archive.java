/** Samples for Reservation Archive. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/Archive.json
     */
    /**
     * Sample code: Archive.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void archive(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .reservations()
            .archiveWithResponse(
                "276e7ae4-84d0-4da6-ab4b-d6b94f3557da",
                "356e7ae4-84d0-4da6-ab4b-d6b94f3557da",
                com.azure.core.util.Context.NONE);
    }
}
