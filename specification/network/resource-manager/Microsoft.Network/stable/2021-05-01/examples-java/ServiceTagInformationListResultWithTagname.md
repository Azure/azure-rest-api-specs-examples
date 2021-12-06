Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ServiceTagInformation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ServiceTagInformationListResultWithTagname.json
     */
    /**
     * Sample code: Get list of service tags with tag name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfServiceTagsWithTagName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getServiceTagInformations()
            .list("westeurope", null, "ApiManagement", Context.NONE);
    }
}
```
