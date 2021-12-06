Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.avs.models.WorkloadNetworkDhcpServer;

/** Samples for WorkloadNetworks CreateDhcp. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateDhcpConfigurations.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateDhcp.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateDhcp(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .defineDhcp("dhcp1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withProperties(
                new WorkloadNetworkDhcpServer()
                    .withDisplayName("dhcpConfigurations1")
                    .withRevision(1L)
                    .withServerAddress("40.1.5.1/24")
                    .withLeaseTime(86400L))
            .create();
    }
}
```
