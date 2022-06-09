```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.ApiVersionSetContract;
import com.azure.resourcemanager.apimanagement.models.VersioningScheme;

/** Samples for ApiVersionSet Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementUpdateApiVersionSet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiVersionSetContract resource =
            manager.apiVersionSets().getWithResponse("rg1", "apimService1", "vs1", Context.NONE).getValue();
        resource
            .update()
            .withDisplayName("api set 1")
            .withVersioningScheme(VersioningScheme.SEGMENT)
            .withDescription("Version configuration")
            .withIfMatch("*")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
