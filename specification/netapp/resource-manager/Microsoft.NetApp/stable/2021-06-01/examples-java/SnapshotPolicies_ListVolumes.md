Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SnapshotPolicies ListVolumes. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/SnapshotPolicies_ListVolumes.json
     */
    /**
     * Sample code: SnapshotPolicies_ListVolumes.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesListVolumes(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().listVolumesWithResponse("myRG", "account1", "snapshotPolicyName", Context.NONE);
    }
}
```
