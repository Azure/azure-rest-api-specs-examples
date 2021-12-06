Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.PrivateCloud;

/** Samples for PrivateClouds Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Update_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_Update_Stretched.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsUpdateStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        PrivateCloud resource =
            manager.privateClouds().getByResourceGroupWithResponse("group1", "cloud1", Context.NONE).getValue();
        resource.update().withManagementCluster(new ManagementCluster().withClusterSize(4)).apply();
    }
}
```
