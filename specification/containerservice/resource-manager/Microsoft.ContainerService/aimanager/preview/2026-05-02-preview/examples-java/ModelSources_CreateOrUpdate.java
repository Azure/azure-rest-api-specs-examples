
import com.azure.resourcemanager.containerserviceaimanager.models.CredentialValue;
import com.azure.resourcemanager.containerserviceaimanager.models.InlineCredential;
import com.azure.resourcemanager.containerserviceaimanager.models.ModelSourceProperties;
import com.azure.resourcemanager.containerserviceaimanager.models.ModelSourceType;

/**
 * Samples for ModelSources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelSources_CreateOrUpdate.json
     */
    /**
     * Sample code: ModelSources_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelSourcesCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelSources().define("huggingface").withExistingAiManager("rgaimanagers", "aimanager1")
            .withProperties(new ModelSourceProperties().withSourceType(ModelSourceType.HUGGING_FACE)
                .withDescription("Hugging Face model source")
                .withCredential(new CredentialValue()
                    .withInline(new InlineCredential().withValue("hf_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"))))
            .withIfMatch("\"00000000-0000-0000-0000-000000000000\"").withIfNoneMatch("*").create();
    }
}
