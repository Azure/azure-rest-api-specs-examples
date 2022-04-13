Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.BlobRestoreParameters;
import com.azure.resourcemanager.storage.models.BlobRestoreRange;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for StorageAccounts RestoreBlobRanges. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobRangesRestore.json
     */
    /**
     * Sample code: BlobRangesRestore.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void blobRangesRestore(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .restoreBlobRanges(
                "res9101",
                "sto4445",
                new BlobRestoreParameters()
                    .withTimeToRestore(OffsetDateTime.parse("2019-04-20T15:30:00.0000000Z"))
                    .withBlobRanges(
                        Arrays
                            .asList(
                                new BlobRestoreRange()
                                    .withStartRange("container/blobpath1")
                                    .withEndRange("container/blobpath2"),
                                new BlobRestoreRange().withStartRange("container2/blobpath3").withEndRange(""))),
                Context.NONE);
    }
}
```
