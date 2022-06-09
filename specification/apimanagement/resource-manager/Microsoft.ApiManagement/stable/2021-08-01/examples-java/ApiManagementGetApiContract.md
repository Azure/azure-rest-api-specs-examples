```java
import com.azure.core.util.Context;

/** Samples for Api Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiContract.json
     */
    /**
     * Sample code: ApiManagementGetApiContract.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiContract(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apis().getWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
