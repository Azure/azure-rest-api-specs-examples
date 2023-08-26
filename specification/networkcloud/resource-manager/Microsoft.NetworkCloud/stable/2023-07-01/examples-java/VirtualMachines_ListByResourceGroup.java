/** Samples for VirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/VirtualMachines_ListByResourceGroup.json
     */
    /**
     * Sample code: List virtual machines for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listVirtualMachinesForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.virtualMachines().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
