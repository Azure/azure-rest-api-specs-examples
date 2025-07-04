
/**
 * Samples for Reservations GetResourceLimits.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Reservations_GetResourceLimits_MaximumSet_Gen.json
     */
    /**
     * Sample code: Reservations_GetResourceLimits.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        reservationsGetResourceLimits(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.reservations().getResourceLimitsWithResponse("rgpurestorage", "storagePoolname",
            com.azure.core.util.Context.NONE);
    }
}
