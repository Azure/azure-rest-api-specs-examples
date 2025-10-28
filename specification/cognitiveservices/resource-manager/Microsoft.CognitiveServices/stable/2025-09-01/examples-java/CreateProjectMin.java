
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ProjectProperties;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;

/**
 * Samples for Projects Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * CreateProjectMin.json
     */
    /**
     * Sample code: Create Project Min.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createProjectMin(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projects().define("testProject1").withExistingAccount("myResourceGroup", "testCreate1")
            .withRegion("West US").withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(new ProjectProperties()).create();
    }
}
