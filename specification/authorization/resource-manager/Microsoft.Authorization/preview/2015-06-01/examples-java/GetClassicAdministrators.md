Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ClassicAdministrators List. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2015-06-01/examples/GetClassicAdministrators.json
     */
    /**
     * Sample code: List classic administrators.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listClassicAdministrators(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getClassicAdministrators()
            .list(Context.NONE);
    }
}
```
