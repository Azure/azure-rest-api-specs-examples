Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BlobServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/BlobServicesList.json
     */
    /**
     * Sample code: ListBlobServices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBlobServices(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobServices().list("res4410", "sto8607", Context.NONE);
    }
}
```
