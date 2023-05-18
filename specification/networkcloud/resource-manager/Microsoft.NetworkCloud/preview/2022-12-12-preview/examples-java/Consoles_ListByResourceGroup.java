/** Samples for Consoles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Consoles_ListByResourceGroup.json
     */
    /**
     * Sample code: List virtual machine consoles for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listVirtualMachineConsolesForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .consoles()
            .listByResourceGroup("resourceGroupName", "virtualMachineName", com.azure.core.util.Context.NONE);
    }
}
