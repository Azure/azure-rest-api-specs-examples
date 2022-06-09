```java
import com.azure.core.util.Context;

/** Samples for Tag GetByOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperationTag.json
     */
    /**
     * Sample code: ApiManagementGetApiOperationTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiOperationTag(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .getByOperationWithResponse(
                "rg1",
                "apimService1",
                "59d6bb8f1f7fab13dc67ec9b",
                "59d6bb8f1f7fab13dc67ec9a",
                "59306a29e4bbd510dc24e5f9",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
