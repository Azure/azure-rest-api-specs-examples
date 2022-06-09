```java
import com.azure.core.util.Context;

/** Samples for Apps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_List.json
     */
    /**
     * Sample code: Apps_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApps().list("myResourceGroup", "myservice", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
