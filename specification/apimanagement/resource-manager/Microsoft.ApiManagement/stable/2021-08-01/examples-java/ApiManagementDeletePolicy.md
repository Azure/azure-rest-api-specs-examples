```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for Policy Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeletePolicy.json
     */
    /**
     * Sample code: ApiManagementDeletePolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeletePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policies().deleteWithResponse("rg1", "apimService1", PolicyIdName.POLICY, "*", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
