Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.6/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.Volume;

/** Samples for Volumes Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-06-01/examples/Volumes_Update.json
     */
    /**
     * Sample code: Volumes_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        Volume resource =
            manager.volumes().getWithResponse("myRG", "account1", "pool1", "volume1", Context.NONE).getValue();
        resource.update().apply();
    }
}
```
