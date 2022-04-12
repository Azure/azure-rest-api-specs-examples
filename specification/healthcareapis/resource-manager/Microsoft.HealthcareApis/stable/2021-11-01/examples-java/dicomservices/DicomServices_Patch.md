Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.healthcareapis.models.DicomService;
import java.util.HashMap;
import java.util.Map;

/** Samples for DicomServices Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Patch.json
     */
    /**
     * Sample code: Update a dicomservice.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void updateADicomservice(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        DicomService resource =
            manager.dicomServices().getWithResponse("testRG", "workspace1", "blue", Context.NONE).getValue();
        resource.update().withTags(mapOf("tagKey", "tagValue")).apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
