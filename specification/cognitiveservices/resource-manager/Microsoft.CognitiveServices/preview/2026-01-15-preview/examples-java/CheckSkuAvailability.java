
import com.azure.resourcemanager.cognitiveservices.models.CheckSkuAvailabilityParameter;
import java.util.Arrays;

/**
 * Samples for ResourceProvider CheckSkuAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/CheckSkuAvailability.json
     */
    /**
     * Sample code: Check SKU Availability.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        checkSKUAvailability(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.resourceProviders()
            .checkSkuAvailabilityWithResponse("westus", new CheckSkuAvailabilityParameter()
                .withSkus(Arrays.asList("S0")).withKind("Face").withType("Microsoft.CognitiveServices/accounts"),
                com.azure.core.util.Context.NONE);
    }
}
