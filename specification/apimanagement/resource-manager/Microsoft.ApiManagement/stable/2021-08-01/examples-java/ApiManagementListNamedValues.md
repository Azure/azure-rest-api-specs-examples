```java
import com.azure.core.util.Context;

/** Samples for NamedValue ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListNamedValues.json
     */
    /**
     * Sample code: ApiManagementListNamedValues.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListNamedValues(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.namedValues().listByService("rg1", "apimService1", null, null, null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
