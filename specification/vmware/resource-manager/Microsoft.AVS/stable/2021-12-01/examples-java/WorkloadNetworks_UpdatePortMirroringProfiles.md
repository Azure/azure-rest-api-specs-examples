Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.PortMirroringDirectionEnum;
import com.azure.resourcemanager.avs.models.WorkloadNetworkPortMirroring;

/** Samples for WorkloadNetworks UpdatePortMirroring. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdatePortMirroringProfiles.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdatePortMirroring.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdatePortMirroring(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkPortMirroring resource =
            manager
                .workloadNetworks()
                .getPortMirroringWithResponse("group1", "cloud1", "portMirroring1", Context.NONE)
                .getValue();
        resource
            .update()
            .withDirection(PortMirroringDirectionEnum.BIDIRECTIONAL)
            .withSource("vmGroup1")
            .withDestination("vmGroup2")
            .withRevision(1L)
            .apply();
    }
}
```
