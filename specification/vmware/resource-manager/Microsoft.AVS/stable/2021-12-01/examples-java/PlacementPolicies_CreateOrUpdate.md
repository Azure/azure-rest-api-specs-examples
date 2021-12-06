Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.avs.models.AffinityType;
import com.azure.resourcemanager.avs.models.VmHostPlacementPolicyProperties;
import java.util.Arrays;

/** Samples for PlacementPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PlacementPolicies_CreateOrUpdate.json
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
                                "/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128",
                                "/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"))
                    .withHostMembers(
                        Arrays
                            .asList(
                                "fakehost22.nyc1.kubernetes.center",
                                "fakehost23.nyc1.kubernetes.center",
                                "fakehost24.nyc1.kubernetes.center"))
                    .withAffinityType(AffinityType.ANTI_AFFINITY))
            .create();
    }
}
```
