Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ApiSchema GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiSchema.json
     */
    /**
     * Sample code: ApiManagementHeadApiSchema.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiSchema(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiSchemas()
            .getEntityTagWithResponse(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                "ec12520d-9d48-4e7b-8f39-698ca2ac63f1",
                Context.NONE);
    }
}
```
