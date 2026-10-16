
import com.azure.resourcemanager.cognitiveservices.models.RaiBindingProperties;

/**
 * Samples for RaiBindings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/PutRaiBinding.json
     */
    /**
     * Sample code: Bind an Azure OpenAI deployment to an ACS policy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void bindAnAzureOpenAIDeploymentToAnACSPolicy(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBindings().define("chat-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiBindingProperties().withBoundResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/aoai-account/deployments/chat-prod")
                .withTargetPolicyName("agent-guard"))
            .withIfNoneMatch("*").create();
    }
}
