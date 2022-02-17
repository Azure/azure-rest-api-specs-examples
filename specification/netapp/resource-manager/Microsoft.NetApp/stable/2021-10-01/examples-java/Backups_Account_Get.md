Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.8/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AccountBackups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Backups_Account_Get.json
     */
    /**
     * Sample code: AccountBackups_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountBackupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accountBackups().getWithResponse("myRG", "account1", "backup1", Context.NONE);
    }
}
```
