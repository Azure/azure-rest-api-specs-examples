Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/PrivateLinkResourcesListByService.json
     */
    /**
     * Sample code: PrivateLinkResources_ListGroupIds.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void privateLinkResourcesListGroupIds(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateLinkResources().listByServiceWithResponse("rgname", "service1", Context.NONE);
    }
}
```
