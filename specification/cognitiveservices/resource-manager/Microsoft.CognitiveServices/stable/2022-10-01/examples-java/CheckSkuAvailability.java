import com.azure.core.util.Context;
import com.azure.resourcemanager.cognitiveservices.models.CheckSkuAvailabilityParameter;
import java.util.Arrays;

/** Samples for ResourceProvider CheckSkuAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/CheckSkuAvailability.json
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
