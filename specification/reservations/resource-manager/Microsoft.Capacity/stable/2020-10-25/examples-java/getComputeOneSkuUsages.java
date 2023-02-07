/** Samples for Quota Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeOneSkuUsages.json
     */
    /**
     * Sample code: Quotas_Request_ForCompute.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void quotasRequestForCompute(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .quotas()
            .getWithResponse(
                "00000000-0000-0000-0000-000000000000",
                "Microsoft.Compute",
                "eastus",
                "standardNDSFamily",
                com.azure.core.util.Context.NONE);
    }
}
