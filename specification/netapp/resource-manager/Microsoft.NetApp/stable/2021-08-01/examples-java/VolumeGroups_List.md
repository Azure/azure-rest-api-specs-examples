Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.7/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VolumeGroups ListByNetAppAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/VolumeGroups_List.json
     */
    /**
     * Sample code: VolumeGroups_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().listByNetAppAccount("myRG", "account1", Context.NONE);
    }
}
```
