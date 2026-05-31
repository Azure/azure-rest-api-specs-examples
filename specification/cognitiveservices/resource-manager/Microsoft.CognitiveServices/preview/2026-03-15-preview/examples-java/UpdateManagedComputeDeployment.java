
import com.azure.resourcemanager.cognitiveservices.models.ManagedComputeDeployment;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for ManagedComputeDeployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/UpdateManagedComputeDeployment.json
     */
    /**
     * Sample code: UpdateManagedComputeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        updateManagedComputeDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        ManagedComputeDeployment resource = manager.managedComputeDeployments()
            .getWithResponse("resourceGroupName", "accountName", "gpt-oss-120b-gpu", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withSku(new Sku().withName("GlobalManagedCompute").withCapacity(2)).apply();
    }
}
