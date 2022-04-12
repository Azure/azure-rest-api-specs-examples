Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FhirServices ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_List.json
     */
    /**
     * Sample code: List fhirservices.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listFhirservices(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirServices().listByWorkspace("testRG", "workspace1", Context.NONE);
    }
}
```
