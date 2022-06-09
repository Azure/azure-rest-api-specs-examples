```java
import com.azure.core.util.Context;

/** Samples for ApiSchema Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiSchema.json
     */
    /**
     * Sample code: ApiManagementGetApiSchema.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiSchemas()
            .getWithResponse(
                "rg1",
                "apimService1",
                "59d6bb8f1f7fab13dc67ec9b",
                "ec12520d-9d48-4e7b-8f39-698ca2ac63f1",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
