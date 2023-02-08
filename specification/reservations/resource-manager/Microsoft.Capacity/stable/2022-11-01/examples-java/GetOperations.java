/** Samples for Operation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetOperations.json
     */
    /**
     * Sample code: GetOperations.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void getOperations(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
