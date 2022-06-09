```java
/** Samples for PortalRevision CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreatePortalRevision.json
     */
    /**
     * Sample code: ApiManagementCreatePortalRevision.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreatePortalRevision(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .portalRevisions()
            .define("20201112101010")
            .withExistingService("rg1", "apimService1")
            .withDescription("portal revision 1")
            .withIsCurrent(true)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
