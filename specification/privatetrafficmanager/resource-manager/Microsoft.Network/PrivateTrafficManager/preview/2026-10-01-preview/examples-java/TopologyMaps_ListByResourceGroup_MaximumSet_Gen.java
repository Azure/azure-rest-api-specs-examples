
/**
 * Samples for TopologyMaps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/TopologyMaps_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: TopologyMaps_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void topologyMapsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.topologyMaps().listByResourceGroup("rgprivateTrafficManager", com.azure.core.util.Context.NONE);
    }
}
