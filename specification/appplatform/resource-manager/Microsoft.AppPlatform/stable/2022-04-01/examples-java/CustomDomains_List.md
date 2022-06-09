```java
import com.azure.core.util.Context;

/** Samples for CustomDomains List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/CustomDomains_List.json
     */
    /**
     * Sample code: CustomDomains_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .list("myResourceGroup", "myservice", "myapp", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
