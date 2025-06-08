
/**
 * Samples for Reservations GetBillingReport.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Reservations_GetBillingReport_MaximumSet_Gen.json
     */
    /**
     * Sample code: Reservations_GetBillingReport_MaximumSet.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void reservationsGetBillingReportMaximumSet(
        com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.reservations().getBillingReportWithResponse("rgpurestorage", "reservationname",
            com.azure.core.util.Context.NONE);
    }
}
