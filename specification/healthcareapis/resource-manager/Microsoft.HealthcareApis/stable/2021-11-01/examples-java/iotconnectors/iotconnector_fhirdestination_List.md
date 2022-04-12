Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FhirDestinations ListByIotConnector. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_List.json
     */
    /**
     * Sample code: List IoT Connectors.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listIoTConnectors(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirDestinations().listByIotConnector("testRG", "workspace1", "blue", Context.NONE);
    }
}
```
