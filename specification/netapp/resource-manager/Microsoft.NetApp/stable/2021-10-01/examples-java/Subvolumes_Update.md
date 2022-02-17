Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.8/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.SubvolumeInfo;

/** Samples for Subvolumes Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Subvolumes_Update.json
     */
    /**
     * Sample code: Subvolumes_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        SubvolumeInfo resource =
            manager
                .subvolumes()
                .getWithResponse("myRG", "account1", "pool1", "volume1", "subvolume1", Context.NONE)
                .getValue();
        resource.update().withPath("/subvolumePath").apply();
    }
}
```
