```java
import com.azure.core.util.Context;

/** Samples for Issue Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetIssue.json
     */
    /**
     * Sample code: ApiManagementGetIssue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetIssue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.issues().getWithResponse("rg1", "apimService1", "57d2ef278aa04f0ad01d6cdc", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
