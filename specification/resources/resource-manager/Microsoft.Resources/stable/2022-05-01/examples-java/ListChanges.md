```java
import com.azure.core.util.Context;

/** Samples for Changes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-05-01/examples/ListChanges.json
     */
    /**
     * Sample code: ListChanges.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listChanges(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .resourceChangeClient()
            .getChanges()
            .list("resourceGroup1", "resourceProvider1", "resourceType1", "resourceName1", null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
