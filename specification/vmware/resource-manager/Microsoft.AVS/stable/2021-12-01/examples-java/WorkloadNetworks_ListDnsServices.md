Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkloadNetworks ListDnsServices. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListDnsServices.json
     */
    /**
     * Sample code: WorkloadNetworks_ListDnsServices.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksListDnsServices(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().listDnsServices("group1", "cloud1", Context.NONE);
    }
}
```
