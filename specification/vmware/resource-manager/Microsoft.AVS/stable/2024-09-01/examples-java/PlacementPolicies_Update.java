
import com.azure.resourcemanager.avs.models.AffinityStrength;
import com.azure.resourcemanager.avs.models.AzureHybridBenefitType;
import com.azure.resourcemanager.avs.models.PlacementPolicy;
import com.azure.resourcemanager.avs.models.PlacementPolicyState;
import java.util.Arrays;

/**
 * Samples for PlacementPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PlacementPolicies_Update.json
     */
    /**
     * Sample code: PlacementPolicies_Update.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void placementPoliciesUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        PlacementPolicy resource = manager.placementPolicies()
            .getWithResponse("group1", "cloud1", "cluster1", "policy1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withState(PlacementPolicyState.DISABLED).withVmMembers(Arrays.asList(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128",
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"))
            .withHostMembers(Arrays.asList("fakehost22.nyc1.kubernetes.center", "fakehost23.nyc1.kubernetes.center",
                "fakehost24.nyc1.kubernetes.center"))
            .withAffinityStrength(AffinityStrength.MUST).withAzureHybridBenefitType(AzureHybridBenefitType.SQL_HOST)
            .apply();
    }
}
