```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.ApiReleaseContract;

/** Samples for ApiRelease Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiRelease.json
     */
    /**
     * Sample code: ApiManagementUpdateApiRelease.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiRelease(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiReleaseContract resource =
            manager.apiReleases().getWithResponse("rg1", "apimService1", "a1", "testrev", Context.NONE).getValue();
        resource
            .update()
            .withApiId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/a1")
            .withNotes("yahooagain")
            .withIfMatch("*")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
