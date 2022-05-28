Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.RegenerateTestKeyRequestPayload;
import com.azure.resourcemanager.appplatform.models.TestKeyType;

/** Samples for Services RegenerateTestKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Services_RegenerateTestKey.json
     */
    /**
     * Sample code: Services_RegenerateTestKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesRegenerateTestKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .regenerateTestKeyWithResponse(
                "myResourceGroup",
                "myservice",
                new RegenerateTestKeyRequestPayload().withKeyType(TestKeyType.PRIMARY),
                Context.NONE);
    }
}
```
