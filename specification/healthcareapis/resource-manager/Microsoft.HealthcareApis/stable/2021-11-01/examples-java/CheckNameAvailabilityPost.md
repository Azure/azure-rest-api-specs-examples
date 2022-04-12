Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.healthcareapis.models.CheckNameAvailabilityParameters;

/** Samples for Services CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/CheckNameAvailabilityPost.json
     */
    /**
     * Sample code: Check name availability.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .services()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityParameters()
                    .withName("serviceName")
                    .withType("Microsoft.HealthcareApis/services"),
                Context.NONE);
    }
}
```
