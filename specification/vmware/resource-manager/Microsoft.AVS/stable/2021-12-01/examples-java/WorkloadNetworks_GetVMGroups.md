Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkloadNetworks GetVMGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetVMGroups.json
     */
    /**
     * Sample code: WorkloadNetworks_GetVMGroup.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void workloadNetworksGetVMGroup(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.workloadNetworks().getVMGroupWithResponse("group1", "cloud1", "vmGroup1", Context.NONE);
    }
}
```
