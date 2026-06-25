
/**
 * Samples for BareMetalMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_ListBySubscription.json
     */
    /**
     * Sample code: List bare metal machines for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listBareMetalMachinesForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().list(null, null, com.azure.core.util.Context.NONE);
    }
}
