
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ProjectProperties;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;

/**
 * Samples for Projects Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/CreateProjectMin.json
     */
    /**
     * Sample code: Create Project Min.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createProjectMin(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().define("testProject1").withExistingAccount("myResourceGroup", "testCreate1")
            .withRegion("West US").withProperties(new ProjectProperties())
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)).create();
    }
}
