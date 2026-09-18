
import com.azure.resourcemanager.containerserviceaimanager.models.CredentialValue;
import com.azure.resourcemanager.containerserviceaimanager.models.ManagedIdentityCredential;
import com.azure.resourcemanager.containerserviceaimanager.models.MicrosoftFoundrySource;
import com.azure.resourcemanager.containerserviceaimanager.models.ModelSourceProperties;
import com.azure.resourcemanager.containerserviceaimanager.models.ModelSourceType;

/**
 * Samples for ModelSources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/ModelSources_CreateOrUpdate_ManagedIdentity.json
     */
    /**
     * Sample code: ModelSources_CreateOrUpdate_ManagedIdentity.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelSourcesCreateOrUpdateManagedIdentity(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelSources().define("foundry").withExistingAiManager("rgaimanagers", "aimanager1")
            .withProperties(new ModelSourceProperties().withSourceType(ModelSourceType.MICROSOFT_FOUNDRY)
                .withDescription("Foundry model source")
                .withCredential(
                    new CredentialValue().withManagedIdentity(new ManagedIdentityCredential().withResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/mytestidentity")))
                .withMicrosoftFoundry(new MicrosoftFoundrySource().withProjectResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.CognitiveServices/accounts/test-account/projects/test-model-project")))
            .withIfMatch("\"00000000-0000-0000-0000-000000000000\"").withIfNoneMatch("*").create();
    }
}
