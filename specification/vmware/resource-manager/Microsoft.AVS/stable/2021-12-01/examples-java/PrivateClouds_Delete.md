Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateClouds Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Delete.json
     */
    /**
     * Sample code: PrivateClouds_Delete.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().delete("group1", "cloud1", Context.NONE);
    }
}
```
