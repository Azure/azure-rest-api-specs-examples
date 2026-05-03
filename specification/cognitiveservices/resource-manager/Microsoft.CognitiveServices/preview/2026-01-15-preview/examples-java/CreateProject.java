
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ProjectProperties;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;

/**
 * Samples for Projects Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/CreateProject.json
     */
    /**
     * Sample code: Create Project.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createProject(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().define("testProject1").withExistingAccount("myResourceGroup", "testCreate1")
            .withRegion("West US")
            .withProperties(
                new ProjectProperties().withDisplayName("p1").withDescription("Description of this project"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)).create();
    }
}
