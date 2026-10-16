
import com.azure.resourcemanager.cognitiveservices.models.RaiBinding;
import com.azure.resourcemanager.cognitiveservices.models.RaiBindingProperties;

/**
 * Samples for RaiBindings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/UpdateRaiBinding.json
     */
    /**
     * Sample code: Update a RAI binding target conditionally across subscriptions.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateARAIBindingTargetConditionallyAcrossSubscriptions(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        RaiBinding resource = manager.raiBindings()
            .getWithResponse("resource-group", "safety-account", "chat-guard", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new RaiBindingProperties().withBoundResourceId(
            "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/target-resource-group/providers/Microsoft.CognitiveServices/accounts/aoai-account/deployments/chat-canary")
            .withTargetPolicyName("agent-guard-v2")).withIfMatch("\"00000000-0000-0000-0000-000000000002\"").apply();
    }
}
