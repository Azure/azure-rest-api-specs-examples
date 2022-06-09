```java
import com.azure.resourcemanager.apimanagement.models.ResourceLocationDataContract;

/** Samples for Gateway CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGateway.json
     */
    /**
     * Sample code: ApiManagementCreateGateway.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGateway(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gateways()
            .define("gw1")
            .withExistingService("rg1", "apimService1")
            .withLocationData(new ResourceLocationDataContract().withName("my location"))
            .withDescription("my gateway 1")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
