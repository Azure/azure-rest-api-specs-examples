```java
import com.azure.core.util.Context;

/** Samples for AccountBackups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Backups_Account_List.json
     */
    /**
     * Sample code: AccountBackups_List.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountBackupsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accountBackups().list("myRG", "account1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.8/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.
