/** Samples for Racks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Racks_ListBySubscription.json
     */
    /**
     * Sample code: List racks for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listRacksForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.racks().list(com.azure.core.util.Context.NONE);
    }
}
