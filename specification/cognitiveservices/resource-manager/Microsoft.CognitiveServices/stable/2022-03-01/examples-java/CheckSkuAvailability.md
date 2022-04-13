Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-cognitiveservices_1.0.0-beta.4/sdk/cognitiveservices/azure-resourcemanager-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cognitiveservices.models.CheckSkuAvailabilityParameter;
import java.util.Arrays;

/** Samples for ResourceProvider CheckSkuAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/CheckSkuAvailability.json
     */
    /**
     * Sample code: Check SKU Availability.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void checkSKUAvailability(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .resourceProviders()
            .checkSkuAvailabilityWithResponse(
                "westus",
                new CheckSkuAvailabilityParameter()
                    .withSkus(Arrays.asList("S0"))
                    .withKind("Face")
                    .withType("Microsoft.CognitiveServices/accounts"),
                Context.NONE);
    }
}
```
