```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.WorkloadNetworkDnsZone;
import java.util.Arrays;

/** Samples for WorkloadNetworks UpdateDnsZone. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDnsZones.json
     */
    /**
     * Sample code: WorkloadNetworks_UpdateDnsZone.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksUpdateDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        WorkloadNetworkDnsZone resource =
            manager.workloadNetworks().getDnsZoneWithResponse("group1", "cloud1", "dnsZone1", Context.NONE).getValue();
        resource
            .update()
            .withDisplayName("dnsZone1")
            .withDomain(Arrays.asList())
            .withDnsServerIps(Arrays.asList("1.1.1.1"))
            .withSourceIp("8.8.8.8")
            .withRevision(1L)
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
