```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;

/** Samples for BlobContainers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersPut.json
     */
    /**
     * Sample code: PutContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .createWithResponse("res3376", "sto328", "container6185", new BlobContainerInner(), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
