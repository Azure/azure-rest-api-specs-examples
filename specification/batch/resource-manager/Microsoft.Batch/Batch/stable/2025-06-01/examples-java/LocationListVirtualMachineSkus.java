
/**
 * Samples for Location ListSupportedVirtualMachineSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/LocationListVirtualMachineSkus.json
     */
    /**
     * Sample code: LocationListVirtualMachineSkus.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void locationListVirtualMachineSkus(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.locations().listSupportedVirtualMachineSkus("japaneast", null, null, com.azure.core.util.Context.NONE);
    }
}
