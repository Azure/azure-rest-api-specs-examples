```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.Permissions;
import com.azure.resourcemanager.storage.models.ServiceSasParameters;
import com.azure.resourcemanager.storage.models.SignedResource;
import java.time.OffsetDateTime;

/** Samples for StorageAccounts ListServiceSas. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountListServiceSAS.json
     */
    /**
     * Sample code: StorageAccountListServiceSAS.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListServiceSAS(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .listServiceSasWithResponse(
                "res7439",
                "sto1299",
                new ServiceSasParameters()
                    .withCanonicalizedResource("/blob/sto1299/music")
                    .withResource(SignedResource.C)
                    .withPermissions(Permissions.L)
                    .withSharedAccessExpiryTime(OffsetDateTime.parse("2017-05-24T11:32:48.8457197Z")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
