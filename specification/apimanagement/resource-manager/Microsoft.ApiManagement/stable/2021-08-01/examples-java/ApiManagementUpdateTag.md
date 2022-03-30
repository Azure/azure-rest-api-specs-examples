Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.TagContract;

/** Samples for Tag Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateTag.json
     */
    /**
     * Sample code: ApiManagementUpdateTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        TagContract resource =
            manager.tags().getWithResponse("rg1", "apimService1", "temptag", Context.NONE).getValue();
        resource.update().withDisplayName("temp tag").withIfMatch("*").apply();
    }
}
```
