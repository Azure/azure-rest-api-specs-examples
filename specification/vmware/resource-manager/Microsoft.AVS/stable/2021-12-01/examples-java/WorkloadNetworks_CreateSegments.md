```java
import com.azure.resourcemanager.avs.models.WorkloadNetworkSegmentSubnet;
import java.util.Arrays;

/** Samples for WorkloadNetworks CreateSegments. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateSegments.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateSegments.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateSegments(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .defineSegments("segment1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("segment1")
            .withConnectedGateway("/infra/tier-1s/gateway")
            .withSubnet(
                new WorkloadNetworkSegmentSubnet()
                    .withDhcpRanges(Arrays.asList("40.20.0.0-40.20.0.1"))
                    .withGatewayAddress("40.20.20.20/16"))
            .withRevision(1L)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
