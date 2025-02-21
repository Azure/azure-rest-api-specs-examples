
/**
 * Samples for VirtualMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/
     * VirtualMachines_ListBySubscription.json
     */
    /**
     * Sample code: List virtual machines for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listVirtualMachinesForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().list(com.azure.core.util.Context.NONE);
    }
}
