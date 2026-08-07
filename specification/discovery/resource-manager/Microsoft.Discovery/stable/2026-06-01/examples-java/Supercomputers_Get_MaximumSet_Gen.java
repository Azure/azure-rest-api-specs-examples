
/**
 * Samples for Supercomputers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Supercomputers_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Supercomputers_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void supercomputersGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.supercomputers().getByResourceGroupWithResponse("rgdiscovery", "871f8fdcf046bf0e2f",
            com.azure.core.util.Context.NONE);
    }
}
