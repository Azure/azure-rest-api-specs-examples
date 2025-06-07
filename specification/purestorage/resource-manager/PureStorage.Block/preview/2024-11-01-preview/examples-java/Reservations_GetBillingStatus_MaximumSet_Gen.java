
/**
 * Samples for Reservations GetBillingStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Reservations_GetBillingStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: Reservations_GetBillingStatus_MaximumSet.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void reservationsGetBillingStatusMaximumSet(
        com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.reservations().getBillingStatusWithResponse("rgpurestorage", "reservationname",
            com.azure.core.util.Context.NONE);
    }
}
