Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-netapp_1.0.0-beta.8/sdk/netapp/azure-resourcemanager-netapp/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.SnapshotRestoreFiles;
import java.util.Arrays;

/** Samples for Snapshots RestoreFiles. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Snapshots_SingleFileRestore.json
     */
    /**
     * Sample code: Snapshots_SingleFileRestore.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsSingleFileRestore(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .snapshots()
            .restoreFiles(
                "myRG",
                "account1",
                "pool1",
                "volume1",
                "snapshot1",
                new SnapshotRestoreFiles().withFilePaths(Arrays.asList("/dir1/customer1.db", "/dir1/customer2.db")),
                Context.NONE);
    }
}
```
