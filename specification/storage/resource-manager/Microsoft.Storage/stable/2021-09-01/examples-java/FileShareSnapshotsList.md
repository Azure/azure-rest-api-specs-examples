Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FileShares List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileShareSnapshotsList.json
     */
    /**
     * Sample code: ListShareSnapshots.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listShareSnapshots(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .list("res9290", "sto1590", null, null, "snapshots", Context.NONE);
    }
}
```
