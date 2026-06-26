
/**
 * Samples for BareMetalMachines ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_ListByResourceGroup.json
     */
    /**
     * Sample code: List bare metal machines for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listBareMetalMachinesForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
