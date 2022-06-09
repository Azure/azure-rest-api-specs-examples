```java
import com.azure.core.util.Context;

/** Samples for FhirServices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Get.json
     */
    /**
     * Sample code: Get a Fhir Service.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getAFhirService(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirServices().getWithResponse("testRG", "workspace1", "fhirservices1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.
