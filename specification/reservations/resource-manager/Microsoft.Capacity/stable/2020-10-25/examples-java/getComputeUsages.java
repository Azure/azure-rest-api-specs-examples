/** Samples for Quota List. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeUsages.json
     */
    /**
     * Sample code: Quotas_listUsagesForCompute.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotasListUsagesForCompute(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .quotas()
            .list(
                "00000000-0000-0000-0000-000000000000",
                "Microsoft.Compute",
                "eastus",
                com.azure.core.util.Context.NONE);
    }
}
