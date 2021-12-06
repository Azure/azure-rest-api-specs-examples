Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BlobServices GetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/BlobServicesGet.json
     */
    /**
     * Sample code: GetBlobServices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getBlobServices(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobServices()
            .getServicePropertiesWithResponse("res4410", "sto8607", Context.NONE);
    }
}
```
