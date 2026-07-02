
import com.azure.resourcemanager.cognitiveservices.models.ContainerInstanceComputeProperties;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.SshSettings;
import com.azure.resourcemanager.cognitiveservices.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Computes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/PutContainerInstanceCompute.json
     */
    /**
     * Sample code: PutContainerInstanceCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putContainerInstanceCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().define("myContainerInstance").withExistingAccount("rgcognitiveservices", "myAccount")
            .withProperties(new ContainerInstanceComputeProperties().withTargetClusterId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster")
                .withImageLink("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest")
                .withIdleTimeBeforeShutdown("PT30M")
                .withSshSettings(new SshSettings().withSshPublicKey("fakeTokenPlaceholder").withAdminEnabled(true)))
            .withRegion("eastus")
            .withIdentity(new Identity().withType(ResourceIdentityType.USER_ASSIGNED).withUserAssignedIdentities(mapOf(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
                new UserAssignedIdentity())))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
