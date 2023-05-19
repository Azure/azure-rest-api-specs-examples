/** Samples for BareMetalMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachines_ListByResourceGroup.json
     */
    /**
     * Sample code: List bare metal machines for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBareMetalMachinesForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
