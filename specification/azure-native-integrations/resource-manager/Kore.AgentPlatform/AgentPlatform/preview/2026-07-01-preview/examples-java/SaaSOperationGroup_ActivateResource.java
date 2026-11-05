
import com.azure.resourcemanager.koreaiartemis.models.ActivateSaaSParameterRequest;

/**
 * Samples for SaaSOperationGroup ActivateResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/SaaSOperationGroup_ActivateResource.json
     */
    /**
     * Sample code: SaaSOperationGroup_ActivateResource.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        saaSOperationGroupActivateResource(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.saaSOperationGroups().activateResource(new ActivateSaaSParameterRequest()
            .withSaasGuid("12345678-1234-1234-1234-123456789012").withPublisherId("kore-ai"),
            com.azure.core.util.Context.NONE);
    }
}
