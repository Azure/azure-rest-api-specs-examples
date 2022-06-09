```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.ApiContract;

/** Samples for Api Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApi.json
     */
    /**
     * Sample code: ApiManagementUpdateApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiContract resource =
            manager.apis().getWithResponse("rg1", "apimService1", "echo-api", Context.NONE).getValue();
        resource
            .update()
            .withDisplayName("Echo API New")
            .withServiceUrl("http://echoapi.cloudapp.net/api2")
            .withPath("newecho")
            .withIfMatch("*")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
