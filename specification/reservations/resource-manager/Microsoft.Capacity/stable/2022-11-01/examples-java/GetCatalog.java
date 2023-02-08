/** Samples for ResourceProvider List. */
public final class Main {
    /*
     * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/GetCatalog.json
     */
    /**
     * Sample code: Catalog.
     *
     * @param manager Entry point to ReservationsManager.
     */
    public static void catalog(com.azure.resourcemanager.reservations.ReservationsManager manager) {
        manager
            .resourceProviders()
            .list(
                "23bc208b-083f-4901-ae85-4f98c0c3b4b6",
                "VirtualMachines",
                "eastus",
                null,
                null,
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
