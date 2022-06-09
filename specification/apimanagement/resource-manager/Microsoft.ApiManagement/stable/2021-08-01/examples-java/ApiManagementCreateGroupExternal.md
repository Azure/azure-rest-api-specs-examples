```java
import com.azure.resourcemanager.apimanagement.models.GroupType;

/** Samples for Group CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGroupExternal.json
     */
    /**
     * Sample code: ApiManagementCreateGroupExternal.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGroupExternal(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groups()
            .define("aadGroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("NewGroup (samiraad.onmicrosoft.com)")
            .withDescription("new group to test")
            .withType(GroupType.EXTERNAL)
            .withExternalId("aad://samiraad.onmicrosoft.com/groups/83cf2753-5831-4675-bc0e-2f8dc067c58d")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
