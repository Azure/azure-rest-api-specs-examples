Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.BlobServicePropertiesInner;
import com.azure.resourcemanager.storage.models.LastAccessTimeTrackingPolicy;
import com.azure.resourcemanager.storage.models.Name;
import java.util.Arrays;

/** Samples for BlobServices SetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/BlobServicesPutLastAccessTimeBasedTracking.json
     */
    /**
     * Sample code: BlobServicesPutLastAccessTimeBasedTracking.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void blobServicesPutLastAccessTimeBasedTracking(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobServices()
            .setServicePropertiesWithResponse(
                "res4410",
                "sto8607",
                new BlobServicePropertiesInner()
                    .withLastAccessTimeTrackingPolicy(
                        new LastAccessTimeTrackingPolicy()
                            .withEnable(true)
                            .withName(Name.ACCESS_TIME_TRACKING)
                            .withTrackingGranularityInDays(1)
                            .withBlobType(Arrays.asList("blockBlob"))),
                Context.NONE);
    }
}
```
