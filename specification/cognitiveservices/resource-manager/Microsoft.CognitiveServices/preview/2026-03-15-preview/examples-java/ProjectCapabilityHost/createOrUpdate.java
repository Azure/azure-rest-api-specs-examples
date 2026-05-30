
import com.azure.resourcemanager.cognitiveservices.models.ProjectCapabilityHostProperties;
import java.util.Arrays;

/**
 * Samples for ProjectCapabilityHosts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ProjectCapabilityHost/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Project CapabilityHost.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createOrUpdateProjectCapabilityHost(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCapabilityHosts().define("capabilityHostName")
            .withExistingProject("test-rg", "account-1", "project-1")
            .withProperties(
                new ProjectCapabilityHostProperties().withAiServicesConnections(Arrays.asList("aoai_connection"))
                    .withVectorStoreConnections(Arrays.asList("acs_connection"))
                    .withStorageConnections(Arrays.asList("blob_connection"))
                    .withThreadStorageConnections(Arrays.asList("aca_connection")))
            .create();
    }
}
