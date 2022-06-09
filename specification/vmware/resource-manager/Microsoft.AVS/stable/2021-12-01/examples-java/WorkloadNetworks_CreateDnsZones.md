```java
import java.util.Arrays;

/** Samples for WorkloadNetworks CreateDnsZone. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateDnsZones.json
     */
    /**
     * Sample code: WorkloadNetworks_CreateDnsZone.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksCreateDnsZone(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .workloadNetworks()
            .defineDnsZone("dnsZone1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withDisplayName("dnsZone1")
            .withDomain(Arrays.asList())
            .withDnsServerIps(Arrays.asList("1.1.1.1"))
            .withSourceIp("8.8.8.8")
            .withRevision(1L)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
