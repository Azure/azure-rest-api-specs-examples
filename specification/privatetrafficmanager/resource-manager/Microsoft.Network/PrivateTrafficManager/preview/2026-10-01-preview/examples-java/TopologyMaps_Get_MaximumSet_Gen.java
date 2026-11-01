
/**
 * Samples for TopologyMaps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/TopologyMaps_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: TopologyMaps_Get_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void topologyMapsGetMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.topologyMaps().getByResourceGroupWithResponse("rgprivateTrafficManager", "myTopologyMap",
            com.azure.core.util.Context.NONE);
    }
}
