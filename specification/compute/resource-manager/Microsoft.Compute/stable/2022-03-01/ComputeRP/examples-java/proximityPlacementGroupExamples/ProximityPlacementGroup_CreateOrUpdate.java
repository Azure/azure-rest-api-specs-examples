import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.ProximityPlacementGroupInner;
import com.azure.resourcemanager.compute.models.ProximityPlacementGroupPropertiesIntent;
import com.azure.resourcemanager.compute.models.ProximityPlacementGroupType;
import java.util.Arrays;

/** Samples for ProximityPlacementGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_CreateOrUpdate.json
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
                    .withZones(Arrays.asList("1"))
                    .withProximityPlacementGroupType(ProximityPlacementGroupType.STANDARD)
                    .withIntent(
                        new ProximityPlacementGroupPropertiesIntent()
                            .withVmSizes(Arrays.asList("Basic_A0", "Basic_A2"))),
                Context.NONE);
    }
}
