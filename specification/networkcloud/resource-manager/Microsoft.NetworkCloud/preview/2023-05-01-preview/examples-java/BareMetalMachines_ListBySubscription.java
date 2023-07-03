/** Samples for BareMetalMachines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachines_ListBySubscription.json
     */
    /**
     * Sample code: List bare metal machines for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBareMetalMachinesForSubscription(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().list(com.azure.core.util.Context.NONE);
    }
}
