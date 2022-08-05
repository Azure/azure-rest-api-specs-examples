import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.ProximityPlacementGroupInner;
import com.azure.resourcemanager.compute.models.ProximityPlacementGroupType;

/** Samples for ProximityPlacementGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/CreateOrUpdateAProximityPlacementGroup.json
     */
    /**
     * Sample code: Create or Update a proximity placement group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAProximityPlacementGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getProximityPlacementGroups()
            .createOrUpdateWithResponse(
                "myResourceGroup",
                "myProximityPlacementGroup",
                new ProximityPlacementGroupInner()
                    .withLocation("westus")
                    .withProximityPlacementGroupType(ProximityPlacementGroupType.STANDARD),
                Context.NONE);
    }
}
