import com.azure.core.util.Context;

/** Samples for PlacementPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PlacementPolicies_List.json
     */
    /**
     * Sample code: PlacementPolicies_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void placementPoliciesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.placementPolicies().list("group1", "cloud1", "cluster1", Context.NONE);
    }
}
