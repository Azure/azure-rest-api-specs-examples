
import com.azure.resourcemanager.cognitiveservices.fluent.models.RaiExternalSafetyProviderSchemaInner;
import com.azure.resourcemanager.cognitiveservices.models.RaiExternalSafetyProviderSchemaProperties;

/**
 * Samples for RaiExternalSafetyProvider CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/PutRaiExternalSafetyProvider.json
     */
    /**
     * Sample code: PutRaiExternalSafetyProvider.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putRaiExternalSafetyProvider(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiExternalSafetyProviders().createOrUpdateWithResponse("safetyProviderName",
            new RaiExternalSafetyProviderSchemaInner().withProperties(new RaiExternalSafetyProviderSchemaProperties()
                .withProviderId("00000000-0000-0000-0000-000000000000").withProviderName("safetyProviderName")
                .withMode("sync").withUrl("https://example.webhook.endpoint").withSecretName("fakeTokenPlaceholder")
                .withManagedIdentity("00000000-0000-0000-0000-000000000000").withKeyVaultUri("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
