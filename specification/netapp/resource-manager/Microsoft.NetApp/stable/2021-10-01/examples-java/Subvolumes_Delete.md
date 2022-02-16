Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.8/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Subvolumes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Subvolumes_Delete.json
     */
    /**
     * Sample code: Subvolumes_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().delete("myRG", "account1", "pool1", "volume1", "subvolume1", Context.NONE);
    }
}
```
