import com.azure.resourcemanager.avs.models.AffinityStrength;
import com.azure.resourcemanager.avs.models.AffinityType;
import com.azure.resourcemanager.avs.models.AzureHybridBenefitType;
import com.azure.resourcemanager.avs.models.VmHostPlacementPolicyProperties;
import java.util.Arrays;

/** Samples for PlacementPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/PlacementPolicies_CreateOrUpdate.json
     */
    /**
     * Sample code: PlacementPolicies_CreateOrUpdate.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void placementPoliciesCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .placementPolicies()
            .define("policy1")
            .withExistingCluster("group1", "cloud1", "cluster1")
            .withProperties(
                new VmHostPlacementPolicyProperties()
                    .withVmMembers(
                        Arrays
                            .asList(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128",
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"))
                    .withHostMembers(
                        Arrays
                            .asList(
                                "fakehost22.nyc1.kubernetes.center",
                                "fakehost23.nyc1.kubernetes.center",
                                "fakehost24.nyc1.kubernetes.center"))
                    .withAffinityType(AffinityType.ANTI_AFFINITY)
                    .withAffinityStrength(AffinityStrength.MUST)
                    .withAzureHybridBenefitType(AzureHybridBenefitType.SQL_HOST))
            .create();
    }
}
