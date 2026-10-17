
import com.azure.resourcemanager.cognitiveservices.models.CostControl;
import com.azure.resourcemanager.cognitiveservices.models.CostControlPatchProperties;

/**
 * Samples for CostControls Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/CostControl/updateMetadata.json
     */
    /**
     * Sample code: Update metadata while preserving legacy rules.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateMetadataWhilePreservingLegacyRules(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        CostControl resource = manager.costControls().getWithResponse("foundry-resource-group", "foundry-account",
            "legacy-budget", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new CostControlPatchProperties().withDisplayName("Renamed legacy budget"))
            .withIfMatch("\"00000000-0000-0000-0000-000000000001\"").apply();
    }
}
